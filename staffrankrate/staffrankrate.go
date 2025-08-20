package staffrankrate

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type StaffRankBonusRate struct {
	ID       int `json:"-"`
	Grade    int `json:"-"`
	Negative int `json:"-"`
	Positive int `json:"-"`
}

var TableName = "staffRankBonusRate"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		grade INTEGER,
		negative INTEGER,
		positive INTEGER	
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *StaffRankBonusRate) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(id, grade, negative, positive) values (?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, c.ID, c.Grade, c.Negative, c.Positive); err != nil {
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

func (r *Repo) GetAll(ctx context.Context) ([]StaffRankBonusRate, error) {
	q := fmt.Sprintf(`
		SELECT
			id, grade, negative, positive
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []StaffRankBonusRate
	for rows.Next() {
		c := StaffRankBonusRate{}
		if err = rows.Scan(&c.ID, &c.Grade, &c.Negative, &c.Positive); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, _ := r.GetAll(ctx)
	if len(all) > 0 {
		slog.Info("staff rank bonus rates already seeded")
		return nil
	}

	for _, v := range StaffRankBonusRates {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
