package clusterbuildcost

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ClusterBuildCost struct {
	ID             int `json:"-"`
	IDX            int `json:"-"`
	Grade          int `json:"-"`
	FOBNumber      int `json:"-"`
	Gmp            int `json:"gmp"`
	ResourceACount int `json:"resource_a_count"`
	ResourceAID    int `json:"resource_a_id"`
	ResourceBCount int `json:"resource_b_count"`
	ResourceBID    int `json:"resource_b_id"`
	TimeMinute     int `json:"time_minute"`
}

var TableName = "clusterBuildCost"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		idx INTEGER,
		grade INTEGER,
		fob_number INTEGER,
		gmp INTEGER,
		resource_a_count INTEGER,
		resource_a_id INTEGER,
		resource_b_count INTEGER,
		resource_b_id INTEGER,
		time_minute INTEGER,
		UNIQUE(idx,grade,fob_number)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ClusterBuildCost) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`
		insert into %s(
			idx,
			grade,
			fob_number,
			gmp,
			resource_a_count,

			resource_a_id,
			resource_b_count,
			resource_b_id,
			time_minute
		) values (?,?,?,?,?, ?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, c.IDX, c.Grade, c.FOBNumber, c.Gmp, c.ResourceACount, c.ResourceAID, c.ResourceBCount, c.ResourceBID, c.TimeMinute); err != nil {
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

func (r *Repo) GetAll(ctx context.Context) ([]ClusterBuildCost, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			idx,
			grade,
			fob_number,
			gmp,
			
			resource_a_count,
			resource_a_id,
			resource_b_count,
			resource_b_id,
			time_minute
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ClusterBuildCost
	for rows.Next() {
		c := ClusterBuildCost{}
		if err = rows.Scan(
			&c.ID,
			&c.IDX,
			&c.Grade,
			&c.FOBNumber,
			&c.Gmp,

			&c.ResourceACount,
			&c.ResourceAID,
			&c.ResourceBCount,
			&c.ResourceBID,
			&c.TimeMinute,
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
		slog.Info("cluster build costs already seeded")
		return nil
	}

	for _, v := range ClusterBuildCosts {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
