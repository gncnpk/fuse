package pfevent

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type PFEvent struct {
	ID      int
	EventID int
}

var TableName = "pfEvent"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER
		);
		CREATE INDEX IF NOT EXISTS pf_event_idx on %s(event_id);`,
		TableName, TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *PFEvent) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(id, event_id) values (?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, c.ID, c.EventID); err != nil {
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

func (r *Repo) GetAll(ctx context.Context) ([]PFEvent, error) {
	q := fmt.Sprintf(`
		SELECT
			id, event_id
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []PFEvent
	for rows.Next() {
		c := PFEvent{}
		if err = rows.Scan(&c.ID, &c.EventID); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, _ := r.GetAll(ctx)
	if len(all) > 0 {
		slog.Info("pf events already seeded")
		return nil
	}

	for _, v := range PFEvents {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
