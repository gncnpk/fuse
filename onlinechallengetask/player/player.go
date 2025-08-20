package onlinechallengetaskplayer

import (
	"context"
	"database/sql"
	"fmt"
)

type OnlineChallengeTaskPlayer struct {
	ID       int
	PlayerID int
	TaskID   int
	Status   int
}

var TableName = "onlineChallengeTaskPlayer"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player_id INTEGER,
		task_id INTEGER,
		status INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *OnlineChallengeTaskPlayer) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		player_id,
		task_id, 
		status
		) values (?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.TaskID,
		c.Status,
	); err != nil {
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]OnlineChallengeTaskPlayer, error) {
	q := fmt.Sprintf(`
		SELECT
			id, player_id, task_id, status
		FROM %s
		WHERE
			player_id = ?
		ORDER BY task_id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []OnlineChallengeTaskPlayer
	for rows.Next() {
		c := OnlineChallengeTaskPlayer{}
		if err = rows.Scan(
			&c.ID,
			&c.PlayerID,
			&c.TaskID,
			&c.Status,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
