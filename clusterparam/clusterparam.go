package clusterparam

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/clustersecurityparam"
	"github.com/unknown321/fuse/localbaseparam"
	"log/slog"
)

type ClusterParam struct {
	ID                   int
	PlatformID           int
	MotherBaseParamID    int
	Build                localbaseparam.LocalBaseParam
	ClusterSecurityParam clustersecurityparam.ClusterSecurityParam
	SoldierRank          int
}

var TableName = "clusterParam"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			platform_id INTEGER,
			mother_base_param_id INTEGER,
			soldier_rank INTEGER,
			defense_level INTEGER,

			non_lethal INTEGER,
			swimsuit_id INTEGER,
			guard_rank INTEGER,
			equipment_grade INTEGER,
			weapon_range INTEGER,

			has_guards INTEGER,
			platforms_built INTEGER,
			mystery_01 INTEGER,
			mystery_02 INTEGER,

			UNIQUE(platform_id, mother_base_param_id)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ClusterParam) (int, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return -1, err
	}

	q := fmt.Sprintf(`INSERT INTO %s(
			platform_id,
			mother_base_param_id,
			soldier_rank,
			defense_level,
			non_lethal,

			swimsuit_id,
			guard_rank,
			equipment_grade,
			weapon_range,
			has_guards,

			platforms_built,
			mystery_01,
			mystery_02

		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?);`,
		TableName)

	if _, err = tx.ExecContext(ctx, q,
		c.PlatformID,
		c.MotherBaseParamID,
		c.SoldierRank,
		c.ClusterSecurityParam.DefenseLevel,
		c.ClusterSecurityParam.NonLethal,

		c.ClusterSecurityParam.SwimsuitID,
		c.ClusterSecurityParam.GuardRank,
		c.ClusterSecurityParam.EquipmentGrade,
		c.ClusterSecurityParam.WeaponRange,
		c.ClusterSecurityParam.HasGuards,

		c.Build.PlatformsBuilt,
		c.Build.Mystery0,
		c.Build.Mystery1,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "motherBaseParamID", c.MotherBaseParamID)
		if err = tx.Rollback(); err != nil {
			return -1, fmt.Errorf("insert rollback failed: %w", err)
		}
		return -1, err
	}

	res := 0
	qq := fmt.Sprintf(`SELECT id FROM %s WHERE mother_base_param_id = ? AND platform_id = ?;`, TableName)
	row := tx.QueryRowContext(ctx, qq, c.MotherBaseParamID, c.PlatformID)
	if err = row.Scan(&res); err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return res, nil
}

func (r *Repo) Update(ctx context.Context, c *ClusterParam) (int, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return -1, err
	}

	q := fmt.Sprintf(`UPDATE %s SET
			platform_id = ?,
			mother_base_param_id = ?,
			soldier_rank = ?,
			defense_level = ?,
			non_lethal = ?,
              
			swimsuit_id = ?,
			guard_rank = ?,
			equipment_grade = ?,
			weapon_range = ?,
			has_guards = ?,
              
			platforms_built = ?,
			mystery_01 = ?,
			mystery_02 = ?

			WHERE id = ?;`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlatformID,
		c.MotherBaseParamID,
		c.SoldierRank,
		c.ClusterSecurityParam.DefenseLevel,
		c.ClusterSecurityParam.NonLethal,

		c.ClusterSecurityParam.SwimsuitID,
		c.ClusterSecurityParam.GuardRank,
		c.ClusterSecurityParam.EquipmentGrade,
		c.ClusterSecurityParam.WeaponRange,
		c.ClusterSecurityParam.HasGuards,

		c.Build.PlatformsBuilt,
		c.Build.Mystery0,
		c.Build.Mystery1,

		c.ID,
	); err != nil {
		slog.Error("update fail", "error", err.Error(), "table", TableName, "motherBaseParamID", c.MotherBaseParamID, "id", c.ID)
		if err = tx.Rollback(); err != nil {
			return -1, fmt.Errorf("insert rollback failed: %w", err)
		}
		return -1, err
	}

	res := 0
	qq := fmt.Sprintf(`SELECT id FROM %s WHERE mother_base_param_id = ? AND platform_id = ?;`, TableName)
	row := tx.QueryRowContext(ctx, qq, c.MotherBaseParamID, c.PlatformID)
	if err = row.Scan(&res); err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return res, nil
}

func (r *Repo) GetByPlatformID(ctx context.Context, motherBaseParamID int, platform_id int) ([]ClusterParam, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			platform_id,
			mother_base_param_id,
			soldier_rank,
			defense_level,

			non_lethal,
			swimsuit_id,
			guard_rank,
			equipment_grade,
			weapon_range,

			has_guards,
			platforms_built,
			mystery_01,
			mystery_02
		FROM %s 
		WHERE 
			mother_base_param_id = ?
			AND platform_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, motherBaseParamID, platform_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ClusterParam
	for rows.Next() {
		c := ClusterParam{}
		if err = rows.Scan(
			&c.ID,
			&c.PlatformID,
			&c.MotherBaseParamID,
			&c.SoldierRank,
			&c.ClusterSecurityParam.DefenseLevel,

			&c.ClusterSecurityParam.NonLethal,
			&c.ClusterSecurityParam.SwimsuitID,
			&c.ClusterSecurityParam.GuardRank,
			&c.ClusterSecurityParam.EquipmentGrade,
			&c.ClusterSecurityParam.WeaponRange,

			&c.ClusterSecurityParam.HasGuards,
			&c.Build.PlatformsBuilt,
			&c.Build.Mystery0,
			&c.Build.Mystery1,
		); err != nil {
			slog.Error("get fail", "error", err.Error())
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Get(ctx context.Context, motherBaseParamID int) ([]ClusterParam, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			platform_id,
			mother_base_param_id,
			soldier_rank,
			defense_level,

			non_lethal,
			swimsuit_id,
			guard_rank,
			equipment_grade,
			weapon_range,

			has_guards,
			platforms_built,
			mystery_01,
			mystery_02
		FROM %s 
		WHERE 
			mother_base_param_id = ?
		ORDER BY platform_id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, motherBaseParamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ClusterParam
	for rows.Next() {
		c := ClusterParam{}
		if err = rows.Scan(
			&c.ID,
			&c.PlatformID,
			&c.MotherBaseParamID,
			&c.SoldierRank,
			&c.ClusterSecurityParam.DefenseLevel,

			&c.ClusterSecurityParam.NonLethal,
			&c.ClusterSecurityParam.SwimsuitID,
			&c.ClusterSecurityParam.GuardRank,
			&c.ClusterSecurityParam.EquipmentGrade,
			&c.ClusterSecurityParam.WeaponRange,

			&c.ClusterSecurityParam.HasGuards,
			&c.Build.PlatformsBuilt,
			&c.Build.Mystery0,
			&c.Build.Mystery1,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
