package challengetask

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/playertask"
	"github.com/unknown321/fuse/tppmessage"
	"strings"
)

type TaskRewardRepo struct {
	db *sql.DB
}

func (r *TaskRewardRepo) TableName() string {
	return "taskReward"
}

func (r *TaskRewardRepo) WithDB(db *sql.DB) *TaskRewardRepo {
	r.db = db
	return r
}

type TaskRewardEntry struct {
	ID     int
	Name   string
	Reward tppmessage.CmdGetChallengeTaskRewardsReward
}

func (r *TaskRewardRepo) CreateSchema(ctx context.Context) error {
	schema := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		taskID      INTEGER PRIMARY KEY,
		name        TEXT,
		bottom_type INTEGER, 
		rate        INTEGER,
		section     INTEGER,
		type        INTEGER,
		value       INTEGER
	);`, r.TableName())
	_, err := r.db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}

	return nil
}

func (r *TaskRewardRepo) Count(ctx context.Context) (int, error) {
	q := fmt.Sprintf("SELECT count(taskID) from %s", r.TableName())
	row := r.db.QueryRowContext(ctx, q)
	res := 0
	if err := row.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *TaskRewardRepo) Seed(ctx context.Context) error {
	query := "INSERT INTO " + r.TableName() + " (taskID, name, bottom_type, rate, section, type, value) values"

	var valueArgs []interface{}
	for _, v := range ChallengeTaskRewardEntries {
		query += "(?,?,?,?,?,?,?),"
		valueArgs = append(valueArgs, v.ID)
		valueArgs = append(valueArgs, v.Name)
		valueArgs = append(valueArgs, v.Reward.BottomType)
		valueArgs = append(valueArgs, v.Reward.Rate)
		valueArgs = append(valueArgs, v.Reward.Section)
		valueArgs = append(valueArgs, v.Reward.Type)
		valueArgs = append(valueArgs, v.Reward.Value)
	}

	query = strings.TrimSuffix(query, ",")

	if _, err := r.db.ExecContext(ctx, query, valueArgs...); err != nil {
		return err
	}

	return nil
}

func (r *TaskRewardRepo) GetByUser(ctx context.Context, userID int) ([]TaskRewardEntry, error) {
	q := fmt.Sprintf(`SELECT 
			t.taskID,
			t.bottom_type,
			t.rate,
			t.section,
			t.type,
			t.value
			FROM %s t JOIN %s p 
			ON t.taskID = p.taskID
			WHERE p.userID = ?`,

		r.TableName(), playertask.TableName,
	)
	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, fmt.Errorf("get task rewards by user: %w", err)
	}

	var res []TaskRewardEntry
	for rows.Next() {
		v := TaskRewardEntry{}
		if err = rows.Scan(&v.ID, &v.Reward.BottomType, &v.Reward.Rate, &v.Reward.Section, &v.Reward.Type, &v.Reward.Value); err != nil {
			return nil, err
		}

		res = append(res, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
