package fobeventreward

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ERewardType uint8

//go:generate go tool stringer -type=ERewardType
const (
	InvalidRewardType ERewardType = iota
	NormalDefense
	NormalSneak
	OneEventTaskSneak
)

type Reward struct {
	ID      int         `json:"-"` // id from `TaskReward` table
	EventID int         `json:"-"`
	Type    ERewardType `json:"-"`

	Reward     int `json:"reward"`
	TaskTypeID int `json:"task_type_id"`
	Threshold  int `json:"threshold"`
}

var TableName = "fobEventReward"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		type INTEGER,
		reward INTEGER,
		task_type_id INTEGER,

		threshold INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *Reward) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`
			insert into %s(
				id,
				event_id,
				type,
				reward,
				task_type_id,

				threshold) values (?,?,?,?,?, ?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.EventID,
		c.Type,
		c.Reward,
		c.TaskTypeID,

		c.Threshold,
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

func (r *Repo) GetAll(ctx context.Context) ([]Reward, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			event_id,
			type,
			reward,
			task_type_id,

			threshold
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Reward
	for rows.Next() {
		c := Reward{}
		if err = rows.Scan(
			&c.ID,
			&c.EventID,
			&c.Type,
			&c.Reward,
			&c.TaskTypeID,

			&c.Threshold,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) GetByType(ctx context.Context, rType ERewardType) ([]Reward, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			event_id,
			type,
			reward,
			task_type_id,

			threshold
		FROM %s
		WHERE type = ?
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, rType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Reward
	for rows.Next() {
		c := Reward{}
		if err = rows.Scan(
			&c.ID,
			&c.EventID,
			&c.Type,
			&c.Reward,
			&c.TaskTypeID,

			&c.Threshold,
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
		slog.Info("fob rewards already seeded")
		return nil
	}

	for _, v := range FOBRewards {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
