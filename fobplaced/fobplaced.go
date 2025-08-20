package fobplaced

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

//go:generate go tool stringer -type=EPlacedType

type EPlacedType uint8

const (
	PlacedTypeInvalid EPlacedType = iota
	MINE
	CAMERA
)

type Placed struct {
	ClusterParamID int
	Type           EPlacedType
	PlacedIndex    int
	PositionX      int
	PositionY      int
	PositionZ      int
	RotationW      int
	RotationX      int
	RotationY      int
	RotationZ      int
	SecurityIDX    int
}

var TableName = "fobPlaced"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			cluster_param_id INTEGER,
			type INTEGER,
			placed_index INTEGER,
			position_x INTEGER,
			position_y INTEGER,
			position_z INTEGER,
			rotation_w INTEGER,
			rotation_x INTEGER,
			rotation_y INTEGER,
			rotation_z INTEGER,
			security_idx INTEGER,
			PRIMARY KEY(cluster_param_id,type,security_idx)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *Placed) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			cluster_param_id,
			type,
			placed_index,
			position_x,
			position_y,

			position_z,
			rotation_w,
			rotation_x,
			rotation_y,
			rotation_z,

			security_idx
		) VALUES (?,?,?,?,?, ?,?,?,?,?, ?)`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ClusterParamID,
		c.Type,
		c.PlacedIndex,
		c.PositionX,
		c.PositionY,

		c.PositionZ,
		c.RotationW,
		c.RotationX,
		c.RotationY,
		c.RotationZ,
		c.SecurityIDX,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "motherbaseID", c.ClusterParamID)
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

func (r *Repo) Get(ctx context.Context, clusterParamID int, securityIDX int) ([]Placed, error) {
	q := fmt.Sprintf(`SELECT 
			cluster_param_id,
			type,
			placed_index,
			position_x,
			position_y,

			position_z,
			rotation_w,
			rotation_x,
			rotation_y,
			rotation_z,

			security_idx
		FROM %s
		WHERE cluster_param_id = ?
			AND security_idx = ?
		ORDER BY placed_index;`,
		TableName)
	rows, err := r.db.QueryContext(ctx, q, clusterParamID, securityIDX)
	if err != nil {
		slog.Error("get fail", "error", err.Error(), "table", TableName, "clusterParamID", clusterParamID, "securityIDX", securityIDX)
		return nil, err
	}

	var res []Placed
	for rows.Next() {
		c := Placed{}
		if err = rows.Scan(
			&c.ClusterParamID,
			&c.Type,
			&c.PlacedIndex,
			&c.PositionX,
			&c.PositionY,

			&c.PositionZ,
			&c.RotationW,
			&c.RotationX,
			&c.RotationY,
			&c.RotationZ,

			&c.SecurityIDX,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}
