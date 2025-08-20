package abolition

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Abolition struct {
	ID        int
	PlayerID  int
	Timestamp int
}

var TableName = "abolition"

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
		timestamp INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, playerID int) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(player_id, timestamp) values (?,unixepoch());`, TableName)
	if _, err = tx.ExecContext(ctx, q, playerID); err != nil {
		slog.Error("add", "error", err.Error(), "table", TableName)
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

func (r *Repo) GetLatest(ctx context.Context, playerID int) (int, error) {
	q := fmt.Sprintf(`
		SELECT 
			COALESCE(MAX(timestamp),0)
		FROM %s
		WHERE 
			player_id = ?`, TableName)
	row := r.db.QueryRowContext(ctx, q, playerID)
	res := 0
	if err := row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *Repo) GetCount(ctx context.Context, playerID int) (int, error) {
	q := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM %s
		WHERE 
			player_id = ?;
	`, TableName)

	row := r.db.QueryRowContext(ctx, q, playerID)
	res := 0
	if err := row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}
