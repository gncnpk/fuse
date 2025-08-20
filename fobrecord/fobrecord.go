package fobrecord

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type FobRecord struct {
	PlayerID    int
	Insurance   int
	Score       int
	ShieldDate  int
	DefenseLose int
	DefenseWin  int
	SneakLose   int
	SneakWin    int
}

var TableName = "fobRecord"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER PRIMARY KEY,
			insurance INTEGER,
			score INTEGER,
			shield_date INTEGER,
			defense_lose INTEGER,
			defense_win INTEGER,
			sneak_lose INTEGER,
			sneak_win INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *FobRecord) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT INTO %s(
			player_id,
			insurance,
			score,
			shield_date,
			defense_lose,
			defense_win,
			sneak_lose,
			sneak_win
		) values (?,?,?,?,?, ?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.Insurance,
		c.Score,
		c.ShieldDate,
		c.DefenseLose,
		c.DefenseWin,
		c.SneakLose,
		c.SneakWin,
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]FobRecord, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			insurance,
			score,
			shield_date,
			defense_lose,
			defense_win,
			sneak_lose,
			sneak_win
		FROM %s 
		WHERE 
			player_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []FobRecord
	for rows.Next() {
		c := FobRecord{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.Insurance,
			&c.Score,
			&c.ShieldDate,
			&c.DefenseLose,
			&c.DefenseWin,
			&c.SneakLose,
			&c.SneakWin,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
