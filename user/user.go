package user

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

// User is the main primitive connecting first party platforms (Steam, PSN, XBOX) to player instances
// ID is used in CMD_REQAUTH_HTTPS
type User struct {
	ID         int
	PlatformID uint64 // steamID, ps, xbox id?
}

var TableName = "user"

type Repo struct {
	db *sql.DB
}

func (ur *Repo) WithDB(db *sql.DB) *Repo {
	ur.db = db
	return ur
}

func (ur *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		platform_id INTEGER NOT NULL UNIQUE
		);`,
		TableName,
	)

	_, err := ur.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (ur *Repo) Add(ctx context.Context, platformID uint64) (int, error) {
	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}

	q := fmt.Sprintf(`insert into %s(platform_id) values (?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, platformID); err != nil {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("insert rollback failed: %w", err)
		}
		return 0, err
	}

	qq := fmt.Sprintf(`select MAX(id) from %s;`, TableName)
	row := tx.QueryRowContext(ctx, qq)
	res := 0
	if err = row.Scan(&res); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return res, nil
}

func (ur *Repo) AddUser(ctx context.Context, u *User) error {
	tx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(id, platform_id) values (?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, u.ID, u.PlatformID); err != nil {
		slog.Error("user add", "error", err.Error())
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

func (ur *Repo) Count(ctx context.Context) (int, error) {
	q := fmt.Sprintf(`SELECT COALESCE(COUNT(id),0) from %s;`, TableName)
	row := ur.db.QueryRowContext(ctx, q)
	res := 0
	if err := row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (ur *Repo) Get(ctx context.Context, platformID uint64) (*User, error) {
	q := fmt.Sprintf(`select id,platform_id from %s where platform_id = ?`, TableName)
	row := ur.db.QueryRowContext(ctx, q, platformID)
	u := &User{}
	if err := row.Scan(&u.ID, &u.PlatformID); err != nil {
		return nil, err
	}

	return u, nil
}
