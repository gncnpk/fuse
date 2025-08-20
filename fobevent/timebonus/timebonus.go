package fobeventtimebonus

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ETimeBonusType uint8

//go:generate go tool stringer -type=ETimeBonusType
const (
	InvalidTimeBonusType ETimeBonusType = iota
	NormalDefense
	NormalSneak
	OneEventTaskSneak
)

type TimeBonus struct {
	ID      int            `json:"-"`
	EventID int            `json:"-"`
	Type    ETimeBonusType `json:"-"`

	SameTimeBonus [8]int `json:"event_sneak_same_time_bonus"`
	BonusMin      int    `json:"event_sneak_clear_point_min"`
	BonusMax      int    `json:"event_sneak_clear_point_max"`
}

var TableName = "fobEventTimeBonus"

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
		same_time_bonus0 INTEGER,
		same_time_bonus1 INTEGER,
		same_time_bonus2 INTEGER,
		same_time_bonus3 INTEGER,
		same_time_bonus4 INTEGER,
		same_time_bonus5 INTEGER,
		same_time_bonus6 INTEGER,
		same_time_bonus7 INTEGER,
		bonus_min INTEGER,
		bonus_max INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *TimeBonus) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		id,
		event_id,
		type,
		same_time_bonus0,
		same_time_bonus1,
		same_time_bonus2,
		same_time_bonus3,
		same_time_bonus4,
		same_time_bonus5,
		same_time_bonus6,
		same_time_bonus7,
		bonus_min,
		bonus_max) values (?,?,?,?,?,?,?,?,?,?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.EventID,
		c.Type,
		c.SameTimeBonus[0],
		c.SameTimeBonus[1],
		c.SameTimeBonus[2],
		c.SameTimeBonus[3],
		c.SameTimeBonus[4],
		c.SameTimeBonus[5],
		c.SameTimeBonus[6],
		c.SameTimeBonus[7],
		c.BonusMin,
		c.BonusMax,
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

func (r *Repo) GetAll(ctx context.Context) ([]TimeBonus, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			event_id,
			type,
			same_time_bonus0,
			same_time_bonus1,
			same_time_bonus2,
			same_time_bonus3,
			same_time_bonus4,
			same_time_bonus5,
			same_time_bonus6,
			same_time_bonus7,
			bonus_min,
			bonus_max
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []TimeBonus
	for rows.Next() {
		c := TimeBonus{}
		if err = rows.Scan(
			&c.ID,
			&c.EventID,
			&c.Type,
			&c.SameTimeBonus[0],
			&c.SameTimeBonus[1],
			&c.SameTimeBonus[2],
			&c.SameTimeBonus[3],
			&c.SameTimeBonus[4],
			&c.SameTimeBonus[5],
			&c.SameTimeBonus[6],
			&c.SameTimeBonus[7],
			&c.BonusMin,
			&c.BonusMax,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) GetByType(ctx context.Context, bonusType ETimeBonusType) ([]TimeBonus, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			event_id,
			type,
			same_time_bonus0,
			same_time_bonus1,
			same_time_bonus2,
			same_time_bonus3,
			same_time_bonus4,
			same_time_bonus5,
			same_time_bonus6,
			same_time_bonus7,
			bonus_min,
			bonus_max
		FROM %s
		WHERE
			type = ?
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, bonusType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []TimeBonus
	for rows.Next() {
		c := TimeBonus{}
		if err = rows.Scan(
			&c.ID,
			&c.EventID,
			&c.Type,
			&c.SameTimeBonus[0],
			&c.SameTimeBonus[1],
			&c.SameTimeBonus[2],
			&c.SameTimeBonus[3],
			&c.SameTimeBonus[4],
			&c.SameTimeBonus[5],
			&c.SameTimeBonus[6],
			&c.SameTimeBonus[7],
			&c.BonusMin,
			&c.BonusMax,
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
		slog.Info("fob event timeBonus already seeded")
		return nil
	}

	for _, v := range FOBTimeBonuses {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
