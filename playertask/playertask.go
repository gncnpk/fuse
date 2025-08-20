package playertask

import (
	"context"
	"database/sql"
	"fmt"
)

type PlayerTask struct {
	PlayerID int
	TaskID   int
}

type Repo struct {
	db *sql.DB
}

var TableName = "playerTask"

func (r *Repo) WithDB(d *sql.DB) *Repo {
	r.db = d
	return r
}

func (r *Repo) InitDB(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			userID INTEGER,
			taskID INTEGER,
			PRIMARY KEY (userID, taskID)
		);`,
		TableName)
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Insert(ctx context.Context, userID uint64, taskID int) error {
	q := fmt.Sprintf(`INSERT INTO %s (userID, taskID) VALUES (?,?);`, TableName)
	if _, err := r.db.ExecContext(ctx, q, userID, taskID); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, userID uint64) ([]int, error) {
	q := fmt.Sprintf("SELECT taskID FROM %s WHERE userID = ?", TableName)
	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int
	for rows.Next() {
		i := 0
		if err = rows.Scan(&i); err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
