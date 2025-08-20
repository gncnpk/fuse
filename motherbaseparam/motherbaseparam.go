package motherbaseparam

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/constructparam"
	"log/slog"
)

type MotherBaseParam struct {
	ID             int // id 2 is used in fob events?
	PlayerID       int
	AreaID         int
	ConstructParam int
	FobIndex       int // always 0?
	PlatformCount  int
	Price          int // always 0?
	SecurityRank   int // TODO calculated how?
}

var TableName = "motherBaseParam"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			player_id INTEGER DEFAULT 0,
			area_id INTEGER DEFAULT 0,
			fob_index INTEGER DEFAULT 0,
			platform_count INTEGER DEFAULT 0,

			price INTEGER DEFAULT 0,
			security_rank INTEGER DEFAULT 0,
			area_code INTEGER DEFAULT 0,
			color INTEGER DEFAULT 0,
			layout_code INTEGER,

			mysterious INTEGER DEFAULT 0,
			
			UNIQUE(player_id, fob_index)
		);
		UPDATE SQLITE_SEQUENCE SET seq = 10 WHERE name = ?;`, // TODO doesn't work
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q, TableName)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *MotherBaseParam) (int, error) {
	ss := constructparam.ConstructParam{}
	if err := ss.FromInt(c.ConstructParam); err != nil {
		return -1, fmt.Errorf("add or update construct param %d: %w", c.ConstructParam, err)
	}

	if ss.ToInt() != c.ConstructParam {
		return -1, fmt.Errorf("construct param pack check fail, input: %d", c.ConstructParam)
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return -1, err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			area_id,
			fob_index,
			platform_count,
			price,

			security_rank,
			area_code,
			color,
			layout_code,
			mysterious
		) values (?,?,?,?,?, ?,?,?,?,?);`,
		TableName)

	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.AreaID,
		c.FobIndex,
		c.PlatformCount,
		c.Price,

		c.SecurityRank,
		ss.AreaCode,
		ss.Color,
		ss.LayoutCode,
		ss.Mysterious,
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "playerID", c.PlayerID)
		if err = tx.Rollback(); err != nil {
			return -1, fmt.Errorf("insert rollback failed: %w", err)
		}
		return -1, err
	}

	res := 0
	qq := fmt.Sprintf(`SELECT id FROM %s where player_id = ? and fob_index = ?;`, TableName)
	row := tx.QueryRowContext(ctx, qq, c.PlayerID, c.FobIndex)
	if err = row.Scan(&res); err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return res, nil
}

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]MotherBaseParam, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			player_id,
			area_id,
			fob_index,
			platform_count,

			price,
			security_rank,
			area_code,
			color,
			layout_code,

			mysterious
		FROM %s 
		WHERE 
			player_id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []MotherBaseParam
	for rows.Next() {
		c := MotherBaseParam{}
		ss := constructparam.ConstructParam{}
		if err = rows.Scan(
			&c.ID,
			&c.PlayerID,
			&c.AreaID,
			&c.FobIndex,
			&c.PlatformCount,

			&c.Price,
			&c.SecurityRank,
			&ss.AreaCode,
			&ss.Color,
			&ss.LayoutCode,

			&ss.Mysterious,
		); err != nil {
			return nil, err
		}

		c.ConstructParam = ss.ToInt()
		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) GetByMotherBaseID(ctx context.Context, motherBaseID int) ([]MotherBaseParam, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			player_id,
			area_id,
			fob_index,
			platform_count,

			price,
			security_rank,
			area_code,
			color,
			layout_code,

			mysterious
		FROM %s 
		WHERE 
			id = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, motherBaseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []MotherBaseParam
	for rows.Next() {
		c := MotherBaseParam{}
		ss := constructparam.ConstructParam{}
		if err = rows.Scan(
			&c.ID,
			&c.PlayerID,
			&c.AreaID,
			&c.FobIndex,
			&c.PlatformCount,

			&c.Price,
			&c.SecurityRank,
			&ss.AreaCode,
			&ss.Color,
			&ss.LayoutCode,

			&ss.Mysterious,
		); err != nil {
			return nil, err
		}

		c.ConstructParam = ss.ToInt()
		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Get(ctx context.Context, playerID int, fobIndex int) ([]MotherBaseParam, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			player_id,
			area_id,
			fob_index,
			platform_count,

			price,
			security_rank,
			area_code,
			color,
			layout_code,

			mysterious
		FROM %s 
		WHERE 
			player_id = ?
			AND fob_index = ?;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID, fobIndex)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []MotherBaseParam
	for rows.Next() {
		c := MotherBaseParam{}
		ss := constructparam.ConstructParam{}
		if err = rows.Scan(
			&c.ID,
			&c.PlayerID,
			&c.AreaID,
			&c.FobIndex,
			&c.PlatformCount,

			&c.Price,
			&c.SecurityRank,
			&ss.AreaCode,
			&ss.Color,
			&ss.LayoutCode,

			&ss.Mysterious,
		); err != nil {
			return nil, err
		}

		c.ConstructParam = ss.ToInt()
		res = append(res, c)
	}

	return res, nil
}
