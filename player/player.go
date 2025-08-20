package player

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/platform"
	"github.com/unknown321/fuse/user"
	"log/slog"
)

type Sneak struct {
	Grade int
	Rank  int
	Score int
}

type Player struct {
	ID         int
	Platform   platform.Platform
	PlatformID uint64 // steamID, psn, xbox id?
	IDX        int

	// these fields must be calculated from other tables
	EspionageLose int
	EspionageWin  int
	FOBGrade      int
	FOBPoint      int
	FOBRank       int // "Espg. Rank" on nameplate
	IsInsurance   int
	LeagueGrade   int
	LeagueRank    int // "PF Rank" on nameplate
	Playtime      int
	Point         int

	Sneak      Sneak // TODO remove, use FOBGrade/Point/Rank instead?
	StaffCount int
}

var TableName = "player"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		idx INTEGER,
		platform INTEGER DEFAULT 0,
		platform_id INTEGER DEFAULT 0,
		espionage_lose INTEGER DEFAULT 0,
		espionage_win INTEGER DEFAULT 0,
		fob_grade INTEGER DEFAULT 0,
		fob_point INTEGER DEFAULT 0,
		fob_rank INTEGER DEFAULT 0,
		is_insurance INTEGER DEFAULT 0,
		league_grade INTEGER DEFAULT 0,
		league_rank INTEGER DEFAULT 0,
		playtime INTEGER DEFAULT 0,
		point INTEGER DEFAULT 0,
		sneak_grade INTEGER DEFAULT 0,
		sneak_rank INTEGER DEFAULT 0,
		sneak_score INTEGER DEFAULT 0,
		staff_count INTEGER DEFAULT 0
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, plat platform.Platform, platformID uint64) (int, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}

	idx, err := r.GetMaxIDX(ctx, plat, platformID)
	if err != nil {
		return 0, fmt.Errorf("get max idx: %w", err)
	}
	idx += 1

	q := fmt.Sprintf(`insert into %s(idx, platform, platform_id) values (?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, idx, plat, platformID); err != nil {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("insert rollback failed: %w", err)
		}
		return 0, err
	}

	qq := fmt.Sprintf(`select MAX(id) from %s;`, TableName)
	row := tx.QueryRowContext(ctx, qq)
	res := 0
	if err = row.Scan(&res); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *Repo) AddPlayer(ctx context.Context, p *Player) error {
	if p.ID == 0 {
		return fmt.Errorf("player id must be > 0")
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		id,
		idx,
		platform,
		platform_id,
		espionage_lose,
		espionage_win,
		fob_grade,
		fob_point,
		fob_rank,
		is_insurance,
		league_grade,
		league_rank,
		playtime,
		point,
		sneak_grade,
		sneak_rank,
		sneak_score,
		staff_count
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		p.ID,
		p.IDX,
		p.Platform,
		p.PlatformID,
		p.EspionageLose,
		p.EspionageWin,
		p.FOBGrade,
		p.FOBPoint,
		p.FOBRank,
		p.IsInsurance,
		p.LeagueGrade,
		p.LeagueRank,
		p.Playtime,
		p.Point,
		p.Sneak.Grade,
		p.Sneak.Rank,
		p.Sneak.Score,
		p.StaffCount,
	); err != nil {
		slog.Error("add player", "error", err.Error())
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

func (r *Repo) Get(ctx context.Context, plat platform.Platform, platformID uint64, idx int) (*Player, error) {
	q := fmt.Sprintf(`
		SELECT 
			id, idx, platform, platform_id, espionage_lose, espionage_win, fob_grade, fob_point, fob_rank, is_insurance, league_grade, league_rank, playtime, point,
			sneak_grade,
			sneak_rank,
			sneak_score,
			staff_count
		FROM %s
		WHERE 
			platform = ? 
			AND platform_id = ?
			AND idx = ?`, TableName)
	row := r.db.QueryRowContext(ctx, q, plat, platformID, idx)
	p := &Player{}
	if err := row.Scan(
		&p.ID,
		&p.IDX,
		&p.Platform,
		&p.PlatformID,
		&p.EspionageLose,
		&p.EspionageWin,
		&p.FOBGrade,
		&p.FOBPoint,
		&p.FOBRank,
		&p.IsInsurance,
		&p.LeagueGrade,
		&p.LeagueRank,
		&p.Playtime,
		&p.Point,
		&p.Sneak.Grade,
		&p.Sneak.Rank,
		&p.Sneak.Score,
		&p.StaffCount,
	); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *Repo) GetMaxIDX(ctx context.Context, plat platform.Platform, platformID uint64) (int, error) {
	q := fmt.Sprintf(`
		SELECT COALESCE(MAX(idx),0)
		FROM %s
		WHERE 
			platform = ?
			AND platform_id = ?;
	`, TableName)

	row := r.db.QueryRowContext(ctx, q, plat, platformID)
	res := 0
	if err := row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *Repo) GetAllByPlatformID(ctx context.Context, plat platform.Platform, platformID uint64) ([]Player, error) {
	q := fmt.Sprintf(`select 
		id, idx, platform, platform_id, espionage_lose, espionage_win, fob_grade, fob_point, fob_rank, is_insurance, league_grade, league_rank, playtime, point,
		sneak_grade,
		sneak_rank,
		sneak_score,
		staff_count
	FROM %s WHERE platform = ? and platform_id = ?`, TableName)
	rows, err := r.db.QueryContext(ctx, q, plat, platformID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Player
	for rows.Next() {
		p := Player{}
		if err = rows.Scan(
			&p.ID,
			&p.IDX,
			&p.Platform,
			&p.PlatformID,
			&p.EspionageLose,
			&p.EspionageWin,
			&p.FOBGrade,
			&p.FOBPoint,
			&p.FOBRank,
			&p.IsInsurance,
			&p.LeagueGrade,
			&p.LeagueRank,
			&p.Playtime,
			&p.Point,
			&p.Sneak.Grade,
			&p.Sneak.Rank,
			&p.Sneak.Score,
			&p.StaffCount,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

func (r *Repo) GetAllByUserID(ctx context.Context, plat platform.Platform, userID int) ([]Player, error) {
	q := fmt.Sprintf(`
	SELECT 
		p.id,
		p.idx,
		p.platform,
		p.platform_id,
		p.espionage_lose,
		p.espionage_win,
		p.fob_grade,
		p.fob_point,
		p.fob_rank,
		p.is_insurance,
		p.league_grade,
		p.league_rank,
		p.playtime,
		p.point,
		p.sneak_grade,
		p.sneak_rank,
		p.sneak_score,
		p.staff_count
	FROM %s p JOIN %s u ON u.platform_id = p.platform_id
	WHERE p.platform = ? 
		AND u.ID = ?`,
		TableName, user.TableName)
	rows, err := r.db.QueryContext(ctx, q, plat, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Player
	for rows.Next() {
		p := Player{}
		if err = rows.Scan(
			&p.ID,
			&p.IDX,
			&p.Platform,
			&p.PlatformID,
			&p.EspionageLose,
			&p.EspionageWin,
			&p.FOBGrade,
			&p.FOBPoint,
			&p.FOBRank,
			&p.IsInsurance,
			&p.LeagueGrade,
			&p.LeagueRank,
			&p.Playtime,
			&p.Point,
			&p.Sneak.Grade,
			&p.Sneak.Rank,
			&p.Sneak.Score,
			&p.StaffCount,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

func (r *Repo) UpdateWith(ctx context.Context, p *Player) error {
	q := fmt.Sprintf(`
		UPDATE %s SET 
			idx = ?,
			platform = ?, 
			platform_id = ?, 
			espionage_lose = ?, 
			espionage_win = ?, 
			fob_grade = ?, 
			fob_point = ?, 
			fob_rank = ?, 
			is_insurance = ?, 
			league_grade = ?, 
		    league_rank = ?,
			playtime = ?, 
			point = ? ,
			sneak_grade = ?,
			sneak_rank = ?,
			sneak_score = ?,
			staff_count = ?
		WHERE id = ?;`, TableName)

	if _, err := r.db.ExecContext(ctx, q,
		p.IDX,
		p.Platform,
		p.PlatformID,
		p.EspionageLose,
		p.EspionageWin,
		p.FOBGrade,
		p.FOBPoint,
		p.FOBRank,
		p.IsInsurance,
		p.LeagueGrade,
		p.LeagueRank,
		p.Playtime,
		p.Point,
		p.Sneak.Grade,
		p.Sneak.Rank,
		p.Sneak.Score,
		p.StaffCount,
		p.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAll(ctx context.Context) ([]Player, error) {
	q := fmt.Sprintf(`select 
		id, idx, platform, platform_id, espionage_lose, espionage_win, fob_grade, fob_point, fob_rank, is_insurance, league_grade, league_rank, playtime, point ,
			sneak_grade,
			sneak_rank,
			sneak_score,
			staff_count
	FROM %s`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Player
	for rows.Next() {
		p := Player{}
		if err = rows.Scan(
			&p.ID,
			&p.IDX,
			&p.Platform,
			&p.PlatformID,
			&p.EspionageLose,
			&p.EspionageWin,
			&p.FOBGrade,
			&p.FOBPoint,
			&p.FOBRank,
			&p.IsInsurance,
			&p.LeagueGrade,
			&p.LeagueRank,
			&p.Playtime,
			&p.Point,
			&p.Sneak.Grade,
			&p.Sneak.Rank,
			&p.Sneak.Score,
			&p.StaffCount,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

func (r *Repo) GetByID(ctx context.Context, plat platform.Platform, playerID int) (Player, error) {
	q := fmt.Sprintf(`
	SELECT 
		p.id,
		p.idx,
		p.platform,
		p.platform_id,
		p.espionage_lose,
		p.espionage_win,
		p.fob_grade,
		p.fob_point,
		p.fob_rank,
		p.is_insurance,
		p.league_grade,
		p.league_rank,
		p.playtime,
		p.point,
		p.sneak_grade,
		p.sneak_rank,
		p.sneak_score,
		p.staff_count
	FROM %s p
	WHERE p.id = ? and p.platform = ?`,
		TableName)
	rows := r.db.QueryRowContext(ctx, q, playerID, plat)

	p := Player{}
	if err := rows.Scan(
		&p.ID,
		&p.IDX,
		&p.Platform,
		&p.PlatformID,
		&p.EspionageLose,
		&p.EspionageWin,
		&p.FOBGrade,
		&p.FOBPoint,
		&p.FOBRank,
		&p.IsInsurance,
		&p.LeagueGrade,
		&p.LeagueRank,
		&p.Playtime,
		&p.Point,
		&p.Sneak.Grade,
		&p.Sneak.Rank,
		&p.Sneak.Score,
		&p.StaffCount,
	); err != nil {
		return p, err
	}

	return p, nil
}
