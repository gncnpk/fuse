package session

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/unknown321/fuse/coder"
	"github.com/unknown321/fuse/platform"
	"log/slog"
	"time"
)

type Session struct {
	UserID        int
	ID            string
	CryptoKey     []byte
	SmartDeviceID string
	InIp          int
	InPort        int
	ExIp          int
	ExPort        int
	Timestamp     int64

	Platform   platform.Platform
	PlatformID uint64 // steamID, psn, xbox
	PlayerID   int
	Coder      coder.Coder
}

func New() (*Session, error) {
	var err error
	key := make([]byte, 16)
	if _, err = rand.Read(key); err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	sessID := make([]byte, 16)
	if _, err = rand.Read(sessID); err != nil {
		return nil, fmt.Errorf("generate session: %w", err)
	}

	sdi := make([]byte, 60)
	if _, err = rand.Read(sdi); err != nil {
		return nil, fmt.Errorf("generate smart device id: %w", err)
	}

	s := &Session{
		ID:            hex.EncodeToString(sessID),
		CryptoKey:     key,
		SmartDeviceID: hex.EncodeToString(sdi),
		Timestamp:     time.Now().Unix(),
		Coder:         coder.Coder{},
		Platform:      platform.Invalid,
	}

	if err = s.Coder.WithKey(key); err != nil {
		return nil, err
	}

	return s, nil
}

var TableName = "session"
var Timeout = 200

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) *Repo {
	r.db = db
	return r
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		user_id INTEGER PRIMARY KEY,
		id STRING,
		crypto_key STRING,
		smart_device_id STRING,
		in_ip INTEGER,
		in_port INTEGER,
		ex_ip INTEGER,
		ex_port INTEGER,
		player_id INTEGER,
		platform INTEGER,
		timestamp INTEGER
		);
		CREATE INDEX IF NOT EXISTS session_id_idx on %s(id);`,
		TableName, TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, session *Session) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`
		INSERT OR REPLACE INTO %s(
			user_id,
			id,
			crypto_key,
			smart_device_id,
			in_ip,

			in_port,
			ex_ip,
			ex_port,
			player_id,
			platform,

			timestamp
		) values (?,?,?,?,?, ?,?,?,?,?, unixepoch());`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		session.UserID,
		session.ID,
		base64.StdEncoding.EncodeToString(session.CryptoKey),
		session.SmartDeviceID,
		session.InIp,
		session.InPort,
		session.ExIp,
		session.ExPort,
		session.PlayerID,
		session.Platform,
	); err != nil {
		slog.Error("add session", "error", err.Error())
		if err = tx.Rollback(); err != nil {
			return fmt.Errorf("insert rollback failed: %w", err)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Remove(ctx context.Context, session *Session) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`DELETE FROM %s WHERE user_id = ?;`, TableName)
	if _, err = tx.ExecContext(ctx, q, session.UserID); err != nil {
		if err = tx.Rollback(); err != nil {
			return fmt.Errorf("remove rollback failed: %w", err)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, sessionKey string) (Session, error) {
	q := fmt.Sprintf(`SELECT user_id, id, crypto_key, smart_device_id, in_ip, in_port, ex_ip, ex_port, player_id, platform, timestamp FROM %s WHERE id = ?`, TableName)
	row := r.db.QueryRowContext(ctx, q, sessionKey)
	s := Session{}
	var key string
	var err error
	if err = row.Scan(
		&s.UserID,
		&s.ID,
		&key,
		&s.SmartDeviceID,
		&s.InIp,
		&s.InPort,
		&s.ExIp,
		&s.ExPort,
		&s.PlayerID,
		&s.Platform,
		&s.Timestamp,
	); err != nil {
		return Session{}, err
	}

	if s.CryptoKey, err = base64.StdEncoding.DecodeString(key); err != nil {
		return Session{}, fmt.Errorf("decode key %s: %w", key, err)
	}

	if err = s.Coder.WithKey(s.CryptoKey); err != nil {
		return Session{}, err
	}

	return s, nil
}

func (r *Repo) GetByPlatformID(ctx context.Context, platformID string) (*Session, error) {
	q := fmt.Sprintf(`
		SELECT 
			user_id,
			id,
			crypto_key,
			smart_device_id,
			in_ip,
			in_port,
			ex_ip,
			ex_port,
			player_id,
			platform,
			timestamp 
		FROM %s s 
		JOIN user u ON u.id = s.user_id
		WHERE u.steam_id = ?`, TableName)
	row := r.db.QueryRowContext(ctx, q, platformID)
	s := &Session{}
	var key string
	var err error
	if err = row.Scan(s.UserID, s.ID, &key, s.SmartDeviceID, s.InIp, s.InPort, s.ExIp, s.ExPort, s.Timestamp); err != nil {
		return nil, err
	}

	if s.CryptoKey, err = base64.StdEncoding.DecodeString(key); err != nil {
		return nil, fmt.Errorf("decode key %s: %w", key, err)
	}

	if err = s.Coder.WithKey(s.CryptoKey); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) SetIP(ctx context.Context, sessionKey string, inIP int, inPort int, exIP int, exPort int) error {
	q := fmt.Sprintf(`UPDATE %s set in_ip = ?, in_port = ?, ex_ip = ?, ex_port = ? where id = ?`, TableName)
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, q, inIP, inPort, exIP, exPort, sessionKey); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) SetPlayerID(ctx context.Context, sessionKey string, playerID int) error {
	q := fmt.Sprintf(`UPDATE %s set player_id = ? where id = ?`, TableName)
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, q, playerID, sessionKey); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) SetTimestamp(ctx context.Context, sessionKey string, timestamp int64) error {
	q := fmt.Sprintf(`UPDATE %s set timestamp = ? where id = ?`, TableName)
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, q, timestamp, sessionKey); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAll(ctx context.Context) ([]Session, error) {
	q := fmt.Sprintf(`SELECT user_id, id, crypto_key, smart_device_id, in_ip, in_port, ex_ip, ex_port, player_id, platform, timestamp FROM %s`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Session
	for rows.Next() {
		s := Session{}
		var key string
		if err = rows.Scan(
			&s.UserID,
			&s.ID,
			&key,
			&s.SmartDeviceID,
			&s.InIp,
			&s.InPort,
			&s.ExIp,
			&s.ExPort,
			&s.PlayerID,
			&s.Platform,
			&s.Timestamp,
		); err != nil {
			return nil, err
		}

		if s.CryptoKey, err = base64.StdEncoding.DecodeString(key); err != nil {
			return nil, fmt.Errorf("decode key %s: %w", key, err)
		}

		if err = s.Coder.WithKey(s.CryptoKey); err != nil {
			return nil, err
		}

		res = append(res, s)
	}

	return res, nil
}
