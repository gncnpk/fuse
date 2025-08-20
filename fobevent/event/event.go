package fobevent

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Event struct {
	ID         int
	Active     bool
	EndDate    int
	DeleteDate int // ??
	Flag       int
}

var TableName = "fobEvent"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		active INTEGER,
		end_date INTEGER,
		delete_date INTEGER,
		flag INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *Event) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
		id,
		active,
		end_date,
		delete_date,
		flag
	) VALUES (?,?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.Active,
		c.EndDate,
		c.DeleteDate,
		c.Flag,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName)
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

func (r *Repo) GetAll(ctx context.Context) ([]Event, error) {
	q := fmt.Sprintf(`
		SELECT
			id, 
			active,
			end_date,
			delete_date,
			flag
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Event
	for rows.Next() {
		c := Event{}
		if err = rows.Scan(
			&c.ID,
			&c.Active,
			&c.EndDate,
			&c.DeleteDate,
			&c.Flag,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) GetActive(ctx context.Context) ([]Event, error) {
	q := fmt.Sprintf(`
		SELECT
			id, 
			active,
			end_date,
			delete_date,
			flag
		FROM %s
		WHERE active != 0
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Event
	for rows.Next() {
		c := Event{}
		if err = rows.Scan(
			&c.ID,
			&c.Active,
			&c.EndDate,
			&c.DeleteDate,
			&c.Flag,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, _ := r.GetAll(ctx)
	if len(all) > 0 {
		slog.Info("fob events already seeded")
		return nil
	}

	for _, v := range FOBEvents {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
