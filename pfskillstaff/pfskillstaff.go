package pfskillstaff

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type PFSkillStaff struct {
	PlayerID                int
	AllStaffNum             int `json:"all_staff_num"`
	Defender1Num            int `json:"defender1_num"`
	Defender2Num            int `json:"defender2_num"`
	Defender3Num            int `json:"defender3_num"`
	InterceptorMissile1Num  int `json:"interceptor_missile1_num"`
	InterceptorMissile2Num  int `json:"interceptor_missile2_num"`
	InterceptorMissile3Num  int `json:"interceptor_missile3_num"`
	LiquidCarbonMissile1Num int `json:"liquid_carbon_missile1_num"`
	LiquidCarbonMissile2Num int `json:"liquid_carbon_missile2_num"`
	LiquidCarbonMissile3Num int `json:"liquid_carbon_missile3_num"`
	Medic1Num               int `json:"medic1_num"`
	Medic2Num               int `json:"medic2_num"`
	Medic3Num               int `json:"medic3_num"`
	Ranger1Num              int `json:"ranger1_num"`
	Ranger2Num              int `json:"ranger2_num"`
	Ranger3Num              int `json:"ranger3_num"`
	Sentry1Num              int `json:"sentry1_num"`
	Sentry2Num              int `json:"sentry2_num"`
	Sentry3Num              int `json:"sentry3_num"`
}

var TableName = "pfSkillStaff"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER PRIMARY KEY,
			all_staff_num INTEGER,
			defender1_num INTEGER,
			defender2_num INTEGER,
			defender3_num INTEGER,
			interceptor_missile1_num INTEGER,
			interceptor_missile2_num INTEGER,
			interceptor_missile3_num INTEGER,
			liquid_carbon_missile1_num INTEGER,
			liquid_carbon_missile2_num INTEGER,
			liquid_carbon_missile3_num INTEGER,
			medic1_num INTEGER,
			medic2_num INTEGER,
			medic3_num INTEGER,
			ranger1_num INTEGER,
			ranger2_num INTEGER,
			ranger3_num INTEGER,
			sentry1_num INTEGER,
			sentry2_num INTEGER,
			sentry3_num INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *PFSkillStaff) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			all_staff_num,
			defender1_num,
			defender2_num,
			defender3_num,
			interceptor_missile1_num,
			interceptor_missile2_num,
			interceptor_missile3_num,
			liquid_carbon_missile1_num,
			liquid_carbon_missile2_num,
			liquid_carbon_missile3_num,
			medic1_num,
			medic2_num,
			medic3_num,
			ranger1_num,
			ranger2_num,
			ranger3_num,
			sentry1_num,
			sentry2_num,
			sentry3_num
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.AllStaffNum,
		c.Defender1Num,
		c.Defender2Num,
		c.Defender3Num,
		c.InterceptorMissile1Num,
		c.InterceptorMissile2Num,
		c.InterceptorMissile3Num,
		c.LiquidCarbonMissile1Num,
		c.LiquidCarbonMissile2Num,
		c.LiquidCarbonMissile3Num,
		c.Medic1Num,
		c.Medic2Num,
		c.Medic3Num,
		c.Ranger1Num,
		c.Ranger2Num,
		c.Ranger3Num,
		c.Sentry1Num,
		c.Sentry2Num,
		c.Sentry3Num,
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
