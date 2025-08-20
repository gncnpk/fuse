package emblem

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Emblem struct {
	PlayerID int
	IDX      int

	BaseColor  int
	FrameColor int
	PositionX  int
	PositionY  int
	Rotate     int
	Scale      int
	TextureTag int
}

var TableName = "emblem"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER,
			idx INTEGER,
			base_color INTEGER,
			frame_color INTEGER,
			position_x INTEGER,

			position_y INTEGER,
			rotate INTEGER,
			scale INTEGER,
			texture_tag INTEGER,

			UNIQUE(player_id, idx)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

// TODO handle removed emblem element
func (r *Repo) AddOrUpdate(ctx context.Context, c *Emblem) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			idx,
			base_color,
			frame_color,
			position_x,

			position_y,
			rotate,
			scale,
			texture_tag

		) values (?,?,?,?,?, ?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.IDX,
		c.BaseColor,
		c.FrameColor,
		c.PositionX,

		c.PositionY,
		c.Rotate,
		c.Scale,
		c.TextureTag,
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]Emblem, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			idx,
			base_color,
			frame_color,
			position_x,

			position_y,
			rotate,
			scale,
			texture_tag
		FROM %s 
		WHERE 
			player_id = ?
		ORDER BY idx;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Emblem
	for rows.Next() {
		c := Emblem{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.IDX,
			&c.BaseColor,
			&c.FrameColor,
			&c.PositionX,

			&c.PositionY,
			&c.Rotate,
			&c.Scale,
			&c.TextureTag,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
