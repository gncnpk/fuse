package onlinechallengetask

import (
	"context"
	"database/sql"
	"fmt"
	onlinechallengetaskplayer "github.com/unknown321/fuse/onlinechallengetask/player"
	"log/slog"
)

type OnlineChallengeTask struct {
	ID               int
	MissionID        int
	RewardBottomType int
	RewardRate       int
	RewardSection    int
	RewardType       int
	RewardValue      int
	TaskTypeID       int
	Threshold        int
	EndDate          int
	Version          int
}

var TableName = "onlineChallengeTask"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mission_id INTEGER,
		reward_bottom_type INTEGER,
		reward_rate INTEGER,
		reward_section INTEGER,
		reward_type INTEGER,
		reward_value INTEGER,
		task_type_id INTEGER,
		threshold INTEGER,
		end_date INTEGER,
		version INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *OnlineChallengeTask) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		id,
		mission_id,
		reward_bottom_type,
		reward_rate,
		reward_section,
		reward_type,
		reward_value,
		task_type_id,
		threshold,
		end_date,
		version) values (?,?,?,?,?, ?,?,?,?,?, ?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.MissionID,
		c.RewardBottomType,
		c.RewardRate,
		c.RewardSection,
		c.RewardType,
		c.RewardValue,
		c.TaskTypeID,
		c.Threshold,
		c.EndDate,
		c.Version,
	); err != nil {
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

func (r *Repo) GetAll(ctx context.Context) ([]OnlineChallengeTask, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			mission_id,
			reward_bottom_type,
			reward_rate,
			reward_section,
			reward_type,
			reward_value,
			task_type_id,
			threshold,
			end_date,
			version
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []OnlineChallengeTask
	for rows.Next() {
		c := OnlineChallengeTask{}
		if err = rows.Scan(
			&c.ID,
			&c.MissionID,
			&c.RewardBottomType,
			&c.RewardRate,
			&c.RewardSection,
			&c.RewardType,
			&c.RewardValue,
			&c.TaskTypeID,
			&c.Threshold,
			&c.EndDate,
			&c.Version,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]OnlineChallengeTask, error) {
	q := fmt.Sprintf(`
		SELECT
			o.id,
			o.mission_id,
			o.reward_bottom_type,
			o.reward_rate,
			o.reward_section,
			o.reward_type,
			o.reward_value,
			o.task_type_id,
			o.threshold,
			o.end_date,
			o.version
		FROM %s o
		JOIN %s pl ON pl.task_id = o.id
		WHERE 
			pl.player_id = ?
			AND pl.status = 0
		ORDER BY o.id;`, TableName, onlinechallengetaskplayer.TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []OnlineChallengeTask
	for rows.Next() {
		c := OnlineChallengeTask{}
		if err = rows.Scan(
			&c.ID,
			&c.MissionID,
			&c.RewardBottomType,
			&c.RewardRate,
			&c.RewardSection,
			&c.RewardType,
			&c.RewardValue,
			&c.TaskTypeID,
			&c.Threshold,
			&c.EndDate,
			&c.Version,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, _ := r.GetAll(ctx)
	if len(all) > 0 {
		slog.Info("fob online challenge task already seeded")
		return nil
	}

	for _, v := range OnlineChallengeTasks {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
