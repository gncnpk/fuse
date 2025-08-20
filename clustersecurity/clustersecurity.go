package clustersecurity

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ClusterSecurity struct {
	ClusterParamID int
	IDX            int
	IsUnique       int

	Antitheft                 int `json:"antitheft"`
	Camera                    int `json:"camera"`
	CautionArea               int `json:"caution_area"`
	Decoy                     int `json:"decoy"`
	IrSensor                  int `json:"ir_sensor"`
	Mine                      int `json:"mine"`
	Soldier                   int `json:"soldier"`
	Uav                       int `json:"uav"`
	VoluntaryCoordCameraCount int `json:"voluntary_coord_camera_count"`
	VoluntaryCoordMineCount   int `json:"voluntary_coord_mine_count"`
}

func CautionAreaToString(c int64) string {
	var bits []byte
	for i := 7; i >= 0; i-- {
		nibble := byte((c >> (i * 4)) & 0xf)
		if i == 32 && nibble == 0 {
			continue
		}
		bits = append(bits, nibble)
	}

	for i := range bits {
		if bits[i] == 0 {
			bits[i] = 45 // -
			continue
		}
		bits[i] += 64
	}

	return string(bits[:])
}

var TableName = "clusterSecurity"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			cluster_param_id INTEGER,
			idx INTEGER,
			is_unique INTEGER,
			antitheft INTEGER,
			camera INTEGER,

			caution_area INTEGER,
			decoy INTEGER,
			ir_sensor INTEGER,
			mine INTEGER,
			soldier INTEGER,

			uav INTEGER,
			voluntary_coord_camera_count INTEGER,
			voluntary_coord_mine_count INTEGER,

			PRIMARY KEY(cluster_param_id, idx)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ClusterSecurity) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT INTO %s(
			cluster_param_id,
			idx,
			is_unique,
			antitheft,
			camera,

			caution_area,
			decoy,
			ir_sensor,
			mine,
			soldier,

			uav,
			voluntary_coord_camera_count,
			voluntary_coord_mine_count
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ClusterParamID,
		c.IDX,
		c.IsUnique,
		c.Antitheft,
		c.Camera,
		c.CautionArea,
		c.Decoy,
		c.IrSensor,
		c.Mine,
		c.Soldier,
		c.Uav,
		c.VoluntaryCoordCameraCount,
		c.VoluntaryCoordMineCount,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "clusterParamID", c.ClusterParamID)
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

func (r *Repo) Update(ctx context.Context, c *ClusterSecurity) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`UPDATE %s SET
			is_unique = ?,
			antitheft = ?,
			camera = ?,

			caution_area = ?,
			decoy = ?,
			ir_sensor = ?,
			mine = ?,
			soldier = ?,

			uav = ?,
			voluntary_coord_camera_count = ?,
			voluntary_coord_mine_count = ?
			WHERE cluster_param_id = ? and idx = ?;`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.IsUnique,
		c.Antitheft,
		c.Camera,
		c.CautionArea,
		c.Decoy,
		c.IrSensor,
		c.Mine,
		c.Soldier,
		c.Uav,
		c.VoluntaryCoordCameraCount,
		c.VoluntaryCoordMineCount,

		c.ClusterParamID,
		c.IDX,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "clusterParamID", c.ClusterParamID)
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

func (r *Repo) Get(ctx context.Context, clusterParamID int, idx int) ([]ClusterSecurity, error) {
	q := fmt.Sprintf(`
		SELECT
			cluster_param_id,
			idx,
			is_unique,
			antitheft,
			camera,

			caution_area,
			decoy,
			ir_sensor,
			mine,
			soldier,

			uav,
			voluntary_coord_camera_count,
			voluntary_coord_mine_count
		FROM %s 
		WHERE 
			cluster_param_id = ?
			AND idx = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, clusterParamID, idx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ClusterSecurity
	for rows.Next() {
		c := ClusterSecurity{}
		if err = rows.Scan(
			&c.ClusterParamID,
			&c.IDX,
			&c.IsUnique,
			&c.Antitheft,
			&c.Camera,

			&c.CautionArea,
			&c.Decoy,
			&c.IrSensor,
			&c.Mine,
			&c.Soldier,

			&c.Uav,
			&c.VoluntaryCoordCameraCount,
			&c.VoluntaryCoordMineCount,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
