package pfranking

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Ranking struct {
	ID       int
	PlayerID int
	Season   int
	IsShort  int
	Finish   int
}

var TableName = "pfRanking"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player_id INTEGER,
		season INTEGER,
		is_short INTEGER,
		finish INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *Ranking) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
		player_id,
		season,
		is_short,
		finish
	) VALUES (?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.Season,
		c.IsShort,
		c.Finish,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName)
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

func (r *Repo) GetCount(ctx context.Context) (int, error) {
	q := fmt.Sprintf(`
		SELECT COUNT(id) FROM %s;`, TableName)
	rows := r.db.QueryRowContext(ctx, q)
	res := 0
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int, seasonID int, shortSeasonID int) ([]Ranking, error) {
	q := fmt.Sprintf(`
		select s.id, s.player_id, s.season, s.is_short, s.finish from %s s join (SELECT COALESCE(MAX(id),0) as id FROM %s WHERE is_short = 0 and season = ? and player_id = ?) b on b.id = s.id
		UNION
		select s.id, s.player_id, s.season, s.is_short, s.finish from %s s join (SELECT COALESCE(MAX(id),0) as id FROM %s WHERE is_short = 1 and season = ? and player_id = ?) b on b.id = s.id
		;
		`, TableName, TableName, TableName, TableName)

	var res []Ranking

	rows, err := r.db.QueryContext(ctx, q,
		seasonID, playerID,
		shortSeasonID, playerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := Ranking{}
		if err = rows.Scan(&c.ID, &c.PlayerID, &c.Season, &c.IsShort, &c.Finish); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}
