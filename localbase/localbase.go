package localbase

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type LocalBase struct {
	PlayerID      int
	MotherBaseNum int
	NamePlateID   int
	PickupOpen    int
	SectionOpen   int
	Param01       int // TODO to localbaseparam
	Param02       int
	Param03       int
	Param04       int
	Param05       int
	Param06       int
	Param07       int
	Time01        int // TODO time left to build a platform?
	Time02        int
	Time03        int
	Time04        int
	Time05        int
	Time06        int
	Time07        int
}

func (l *LocalBase) WithTime(t []int) error {
	if len(t) != 7 {
		return fmt.Errorf("invalid time length, have %d, want 7", len(t))
	}

	l.Time01 = t[0]
	l.Time02 = t[1]
	l.Time03 = t[2]
	l.Time04 = t[3]
	l.Time05 = t[4]
	l.Time06 = t[5]
	l.Time07 = t[6]

	return nil
}

func (l *LocalBase) WithParam(t []int) error {
	if len(t) != 7 {
		return fmt.Errorf("invalid param length, have %d, want 7", len(t))
	}

	l.Param01 = t[0]
	l.Param02 = t[1]
	l.Param03 = t[2]
	l.Param04 = t[3]
	l.Param05 = t[4]
	l.Param06 = t[5]
	l.Param07 = t[6]

	return nil
}

var TableName = "localBase"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		player_id INTEGER PRIMARY KEY,
		mother_base_num INTEGER,
		name_plate_id INTEGER,
		pickup_open INTEGER,
		section_open INTEGER,
		param_01 INTEGER,
		param_02 INTEGER,
		param_03 INTEGER,
		param_04 INTEGER,
		param_05 INTEGER,
		param_06 INTEGER,
		param_07 INTEGER,
		time_01 INTEGER,
		time_02 INTEGER,
		time_03 INTEGER,
		time_04 INTEGER,
		time_05 INTEGER,
		time_06 INTEGER,
		time_07 INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *LocalBase) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			mother_base_num,
			name_plate_id,
			pickup_open,
			section_open,
			param_01,
			param_02,
			param_03,
			param_04,
			param_05,
			param_06,
			param_07,
			time_01,
			time_02,
			time_03,
			time_04,
			time_05,
			time_06,
			time_07
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.MotherBaseNum,
		c.NamePlateID,
		c.PickupOpen,
		c.SectionOpen,
		c.Param01,
		c.Param02,
		c.Param03,
		c.Param04,
		c.Param05,
		c.Param06,
		c.Param07,
		c.Time01,
		c.Time02,
		c.Time03,
		c.Time04,
		c.Time05,
		c.Time06,
		c.Time07,
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
