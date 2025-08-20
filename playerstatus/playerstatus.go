package playerstatus

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type PlayerStatus struct {
	PlayerID          int
	CumulativeGrade   int
	Gmp               int
	Hero              int
	IsForceBalance    int
	IsWallet          int
	ServerGmp         int // max 995 000 000
	InjuryGmp         int
	InsuranceGmp      int
	LoadoutGmp        int
	MbCoin            int
	SecurityChallenge int
	EspionagePoint    int

	EspionageRatingGrade                int
	FobDeployToSupportersEmergencyCount int
	FobSupportingUserCount              int
	PfRatingDefenseForce                int
	PfRatingDefenseLife                 int
	PfRatingOffenceForce                int
	PfRatingOffenceLife                 int
	PfRatingRank                        int
	TotalDevelopmentGrade               int
	TotalFobSecurityLevel               int

	NamePlateID int // from 0 (empty)  to 52 (gold horse)
}

var TableName = "playerStatus"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER PRIMARY KEY,
			cumulative_grade INTEGER,
			gmp INTEGER,
			hero INTEGER,
			is_force_balance INTEGER,
			is_wallet INTEGER,
			server_gmp INTEGER,
			injury_gmp INTEGER,
			insurance_gmp INTEGER,
			loadout_gmp INTEGER,
			mb_coin INTEGER,
			security_challenge INTEGER,
			espionage_point INTEGER,
			espionage_rating_grade INTEGER,
			fob_deploy_to_supporters_emergency_count INTEGER,
			fob_supporting_user_count INTEGER,
			pf_rating_defense_force INTEGER,
			pf_rating_defense_life INTEGER,
			pf_rating_offence_force INTEGER,
			pf_rating_offence_life INTEGER,
			pf_rating_rank INTEGER,
			total_development_grade INTEGER,
			total_fob_security_level INTEGER,
			nameplate_id INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *PlayerStatus) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			cumulative_grade,
			gmp,
			hero,
			is_force_balance,
			is_wallet,
			server_gmp,
			injury_gmp,
			insurance_gmp,
			loadout_gmp,
			mb_coin,
			security_challenge,
			espionage_point,
			espionage_rating_grade,
			fob_deploy_to_supporters_emergency_count,
			fob_supporting_user_count,
			pf_rating_defense_force,
			pf_rating_defense_life,
			pf_rating_offence_force,
			pf_rating_offence_life,
			pf_rating_rank,
			total_development_grade,
			total_fob_security_level,
			nameplate_id
		) VALUES (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?)`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.CumulativeGrade,
		c.Gmp,
		c.Hero,
		c.IsForceBalance,
		c.IsWallet,
		c.ServerGmp,
		c.InjuryGmp,
		c.InsuranceGmp,
		c.LoadoutGmp,
		c.MbCoin,
		c.SecurityChallenge,
		c.EspionagePoint,
		c.EspionageRatingGrade,
		c.FobDeployToSupportersEmergencyCount,
		c.FobSupportingUserCount,
		c.PfRatingDefenseForce,
		c.PfRatingDefenseLife,
		c.PfRatingOffenceForce,
		c.PfRatingOffenceLife,
		c.PfRatingRank,
		c.TotalDevelopmentGrade,
		c.TotalFobSecurityLevel,
		c.NamePlateID,
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

func (r *Repo) Get(ctx context.Context, playerID int) (PlayerStatus, error) {
	q := fmt.Sprintf(`SELECT 
			player_id,
			cumulative_grade,
			gmp,
			hero,
			is_force_balance,
			is_wallet,
			server_gmp,
			injury_gmp,
			insurance_gmp,
			loadout_gmp,
			mb_coin,
			security_challenge,
			espionage_point,
			espionage_rating_grade,
			fob_deploy_to_supporters_emergency_count,
			fob_supporting_user_count,
			pf_rating_defense_force,
			pf_rating_defense_life,
			pf_rating_offence_force,
			pf_rating_offence_life,
			pf_rating_rank,
			total_development_grade,
			total_fob_security_level,
			nameplate_id
		FROM %s
		WHERE player_id = ?;`,
		TableName)
	rows := r.db.QueryRowContext(ctx, q, playerID)

	c := PlayerStatus{}
	if err := rows.Scan(
		&c.PlayerID,
		&c.CumulativeGrade,
		&c.Gmp,
		&c.Hero,
		&c.IsForceBalance,
		&c.IsWallet,
		&c.ServerGmp,
		&c.InjuryGmp,
		&c.InsuranceGmp,
		&c.LoadoutGmp,
		&c.MbCoin,
		&c.SecurityChallenge,
		&c.EspionagePoint,
		&c.EspionageRatingGrade,
		&c.FobDeployToSupportersEmergencyCount,
		&c.FobSupportingUserCount,
		&c.PfRatingDefenseForce,
		&c.PfRatingDefenseLife,
		&c.PfRatingOffenceForce,
		&c.PfRatingOffenceLife,
		&c.PfRatingRank,
		&c.TotalDevelopmentGrade,
		&c.TotalFobSecurityLevel,
		&c.NamePlateID,
	); err != nil {
		return c, err
	}

	return c, nil
}
