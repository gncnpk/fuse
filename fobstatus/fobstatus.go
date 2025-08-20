package fobstatus

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type FobStatus struct {
	PlayerID  int
	IsRescue  int
	IsReward  int
	SneakMode int
}

var TableName = "fobStatus"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER PRIMARY KEY,
			is_rescue INTEGER,
			is_reward INTEGER,
			sneak_mode INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *FobStatus) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT INTO %s(
			player_id,
			is_rescue,
			is_reward,
			sneak_mode
		) values (?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.IsRescue,
		c.IsReward,
		c.SneakMode,
	); err != nil {
		slog.Error("add failed", "table", TableName, "error", err.Error())
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]FobStatus, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			is_rescue,
			is_reward,
			sneak_mode
		FROM %s 
		WHERE 
			player_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []FobStatus
	for rows.Next() {
		c := FobStatus{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.IsRescue,
			&c.IsReward,
			&c.SneakMode,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) SetRescue(ctx context.Context, playerID int) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`UPDATE %s SET
			is_rescue = 1,
			sneak_mode = 1 
			WHERE player_id = ?;`, TableName)
	if _, err = tx.ExecContext(ctx, q, playerID); err != nil {
		slog.Error("set rescue", "table", TableName, "error", err.Error())
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
