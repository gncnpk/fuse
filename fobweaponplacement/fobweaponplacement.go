package fobweaponplacement

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type WeaponPlacement struct {
	MotherBaseID       int
	EmplacementGunEast int
	EmplacementGunWest int
	GatlingGun         int
	GatlingGunEast     int
	GatlingGunWest     int
	MortarNormal       int
}

var TableName = "fobWeaponPlacement"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			mother_base_id INTEGER PRIMARY KEY,
			emplacement_gun_east INTEGER,
			emplacement_gun_west INTEGER,
			gatling_gun INTEGER,
			gatling_gun_east INTEGER,
			gatling_gun_west INTEGER,
			mortar_normal INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *WeaponPlacement) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT INTO %s(
			mother_base_id,
			emplacement_gun_east,
			emplacement_gun_west,
			gatling_gun,
			gatling_gun_east,

			gatling_gun_west,
			mortar_normal
		) values (?,?,?,?,?, ?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.MotherBaseID,
		c.EmplacementGunEast,
		c.EmplacementGunWest,
		c.GatlingGun,
		c.GatlingGunEast,
		c.GatlingGunWest,
		c.MortarNormal,
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

func (r *Repo) Get(ctx context.Context, motherBaseID int) ([]WeaponPlacement, error) {
	q := fmt.Sprintf(`
		SELECT
			mother_base_id,
			emplacement_gun_east,
			emplacement_gun_west,
			gatling_gun,
			gatling_gun_east,

			gatling_gun_west,
			mortar_normal
		FROM %s 
		WHERE 
			mother_base_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, motherBaseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []WeaponPlacement
	for rows.Next() {
		c := WeaponPlacement{}
		if err = rows.Scan(
			&c.MotherBaseID,
			&c.EmplacementGunEast,
			&c.EmplacementGunWest,
			&c.GatlingGun,
			&c.GatlingGunEast,
			&c.GatlingGunWest,
			&c.MortarNormal,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
