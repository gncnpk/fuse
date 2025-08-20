package sectionstat

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ESection uint8

//go:generate go tool stringer -type=ESection

const (
	Base ESection = iota
	Combat
	Develop
	Medical
	Security
	Spy
	Support // "suport" in json
)

type SectionStat struct {
	PlayerID   int
	SectionID  ESection
	Level      int
	SoldierNum int
}

var TableName = "sectionStat"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER,
			section_id INTEGER,
			level INTEGER,
			soldier_num INTEGER,
			PRIMARY KEY(player_id, section_id)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *SectionStat) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			section_id,
			level,
			soldier_num
		) values (?,?,?,?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.SectionID,
		c.Level,
		c.SoldierNum,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "playerID", c.PlayerID)
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]SectionStat, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			section_id,
			level,
			soldier_num
		FROM %s 
		WHERE 
			player_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []SectionStat
	for rows.Next() {
		c := SectionStat{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.SectionID,
			&c.Level,
			&c.SoldierNum,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) GetBySectionID(ctx context.Context, playerID int, sectionID ESection) ([]SectionStat, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			section_id,
			level,
			soldier_num
		FROM %s 
		WHERE 
			player_id = ? AND section_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID, sectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []SectionStat
	for rows.Next() {
		c := SectionStat{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.SectionID,
			&c.Level,
			&c.SoldierNum,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
