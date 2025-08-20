package pfseason

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Season struct {
	ID        int
	IsShort   int
	PlayerNum int
}

var TableName = "pfSeason"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER,
		is_short INTEGER,
		player_num INTEGER,
		UNIQUE(id, is_short)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *Season) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
		id,
		is_short,
		player_num
	) VALUES (?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.IsShort,
		c.PlayerNum,
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

func (r *Repo) GetCount(ctx context.Context) (int, error) {
	q := fmt.Sprintf(`
		SELECT COUNT(id)
		FROM %s;`, TableName)
	rows := r.db.QueryRowContext(ctx, q)
	res := 0
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	c, err := r.GetCount(ctx)
	if err != nil {
		return err
	}

	if c > 0 {
		slog.Info("pf season already seeded")
		return nil
	}

	if err = r.Add(ctx, &Season{}); err != nil {
		return err
	}
	if err = r.Add(ctx, &Season{IsShort: 1}); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetActive(ctx context.Context) ([]Season, error) {
	q := fmt.Sprintf(`
		SELECT MAX(id), is_short, player_num FROM %s WHERE is_short = 0
		UNION
		SELECT MAX(id), is_short, player_num from %s WHERE is_short = 1;
		`, TableName, TableName)

	var res []Season

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := Season{}
		if err = rows.Scan(&c.ID, &c.IsShort, &c.PlayerNum); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}
