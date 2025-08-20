package serverstatus

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

// TODO rewrite, remove pf

type Status struct {
	PfSeason      int
	PfShortSeason int
}

var TableName = "serverStatus"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			pf_season INTEGER,
			pf_short_season INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Seed(ctx context.Context) error {
	q := fmt.Sprintf(`
		SELECT count(*)
		FROM %s;`, TableName)
	row := r.db.QueryRowContext(ctx, q)
	res := 0
	if err := row.Scan(&res); err != nil {
		return err
	}

	if res > 0 {
		slog.Info("server status already seeded")
		return nil
	}

	if err := r.AddOrUpdate(ctx, &Status{}); err != nil {
		return err
	}

	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *Status) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			pf_season,
			pf_short_season
		) VALUES (?,?)`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PfSeason,
		c.PfShortSeason,
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

func (r *Repo) Get(ctx context.Context, playerID int) ([]Status, error) {
	q := fmt.Sprintf(`SELECT 
			pf_season,
			pf_short_season
		FROM %s;`,
		TableName)
	rows, err := r.db.QueryContext(ctx, q,
		playerID)
	if err != nil {
		slog.Error("get fail", "error", err.Error(), "table", TableName, "playerID", playerID)
		return nil, err
	}

	var res []Status
	for rows.Next() {
		c := Status{}
		if err = rows.Scan(
			&c.PfSeason,
			&c.PfShortSeason,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}
