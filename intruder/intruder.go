package intruder

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Intruder struct {
	PlayerID             int
	OwnerID              int
	MotherBaseID         int
	MotherBasePlatformID int
	Mode                 int // visit, sham etc.
	IsSneak              int
	Timestamp            int64
}

var TableName = "intruder"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER PRIMARY KEY,
			owner_id INTEGER,
			mother_base_id INTEGER,
			platform_id INTEGER,
			mode INTEGER,

			is_sneak INTEGER,
			timestamp INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *Intruder) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			owner_id,
			mother_base_id,
			platform_id,
			mode,

			is_sneak,
			timestamp
		) values (?,?,?,?,?,?,unixepoch());`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.OwnerID,
		c.MotherBaseID,
		c.MotherBasePlatformID,
		c.Mode,
		c.IsSneak,
	); err != nil {
		slog.Error("add failed", "table", TableName, "error", err.Error())
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

func (r *Repo) GetByOwnerID(ctx context.Context, ownerID int) ([]Intruder, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			owner_id,
			mother_base_id,
			platform_id,
			mode,

			is_sneak,
			timestamp
		FROM %s 
		WHERE 
			owner_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Intruder
	for rows.Next() {
		c := Intruder{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.OwnerID,
			&c.MotherBaseID,
			&c.MotherBasePlatformID,
			&c.Mode,

			&c.IsSneak,
			&c.Timestamp,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
