package serverproductparam

import (
	"context"
	"database/sql"
	"fmt"
	serverproductparamplayer "github.com/unknown321/fuse/serverproductparam/player"
	"log/slog"
)

type ServerProductParam struct {
	DevCoin            int `json:"dev_coin"`
	DevGmp             int `json:"dev_gmp"`
	DevItem1           int `json:"dev_item_1"`
	DevItem2           int `json:"dev_item_2"`
	DevPlatlv01        int `json:"dev_platlv01"`
	DevPlatlv02        int `json:"dev_platlv02"`
	DevPlatlv03        int `json:"dev_platlv03"`
	DevPlatlv04        int `json:"dev_platlv04"`
	DevPlatlv05        int `json:"dev_platlv05"`
	DevPlatlv06        int `json:"dev_platlv06"`
	DevPlatlv07        int `json:"dev_platlv07"`
	DevRescount01Value int `json:"dev_rescount01_value"`
	DevRescount02Value int `json:"dev_rescount02_value"`
	DevResource01Id    int `json:"dev_resource01_id"`
	DevResource02Id    int `json:"dev_resource02_id"`
	DevSkil            int `json:"dev_skil"`
	DevSpecial         int `json:"dev_special"`
	DevTime            int `json:"dev_time"`
	ID                 int `json:"id"`
	Type               int `json:"type"`
	UseGmp             int `json:"use_gmp"`
	UseRescount01Value int `json:"use_rescount01_value"`
	UseRescount02Value int `json:"use_rescount02_value"`
	UseResource01Id    int `json:"use_resource01_id"`
	UseResource02Id    int `json:"use_resource02_id"`
	Open               int `json:"-"`
}

var TableName = "serverProductParam"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY,
			dev_coin INTEGER,
			dev_gmp INTEGER,
			dev_item_1 INTEGER,
			dev_item_2 INTEGER,
			dev_platlv01 INTEGER,
			dev_platlv02 INTEGER,
			dev_platlv03 INTEGER,
			dev_platlv04 INTEGER,
			dev_platlv05 INTEGER,
			dev_platlv06 INTEGER,
			dev_platlv07 INTEGER,
			dev_rescount01_value INTEGER,
			dev_rescount02_value INTEGER,
			dev_resource01_id INTEGER,
			dev_resource02_id INTEGER,
			dev_skil INTEGER,
			dev_special INTEGER,
			dev_time INTEGER,
			type INTEGER,
			use_gmp INTEGER,
			use_rescount01_value INTEGER,
			use_rescount02_value INTEGER,
			use_resource01_id INTEGER,
			use_resource02_id INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ServerProductParam) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
			id,
			dev_coin,
			dev_gmp,
			dev_item_1,
			dev_item_2,
			dev_platlv01,
			dev_platlv02,
			dev_platlv03,
			dev_platlv04,
			dev_platlv05,
			dev_platlv06,
			dev_platlv07,
			dev_rescount01_value,
			dev_rescount02_value,
			dev_resource01_id,
			dev_resource02_id,
			dev_skil,
			dev_special,
			dev_time,
			type,
			use_gmp,
			use_rescount01_value,
			use_rescount02_value,
			use_resource01_id,
			use_resource02_id
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.DevCoin,
		c.DevGmp,
		c.DevItem1,
		c.DevItem2,
		c.DevPlatlv01,
		c.DevPlatlv02,
		c.DevPlatlv03,
		c.DevPlatlv04,
		c.DevPlatlv05,
		c.DevPlatlv06,
		c.DevPlatlv07,
		c.DevRescount01Value,
		c.DevRescount02Value,
		c.DevResource01Id,
		c.DevResource02Id,
		c.DevSkil,
		c.DevSpecial,
		c.DevTime,
		c.Type,
		c.UseGmp,
		c.UseRescount01Value,
		c.UseRescount02Value,
		c.UseResource01Id,
		c.UseResource02Id,
		c.Open,
	); err != nil {
		slog.Error("tx err", "error", err.Error())
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
		SELECT
			COUNT(id)
		FROM %s;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var res int
	for rows.Next() {
		if err = rows.Scan(
			&res,
		); err != nil {
			return 0, err
		}
	}

	return res, nil
}
func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]ServerProductParam, error) {
	q := fmt.Sprintf(`
		SELECT
			o.id,
			o.dev_coin,
			o.dev_gmp,
			o.dev_item_1,
			o.dev_item_2,
			o.dev_platlv01,
			o.dev_platlv02,
			o.dev_platlv03,
			o.dev_platlv04,
			o.dev_platlv05,
			o.dev_platlv06,
			o.dev_platlv07,
			o.dev_rescount01_value,
			o.dev_rescount02_value,
			o.dev_resource01_id,
			o.dev_resource02_id,
			o.dev_skil,
			o.dev_special,
			o.dev_time,
			o.type,
			o.use_gmp,
			o.use_rescount01_value,
			o.use_rescount02_value,
			o.use_resource01_id,
			o.use_resource02_id,
			pl.open
		FROM %s o
		JOIN %s pl ON pl.product_id = o.id
		WHERE 
			pl.player_id = ?
		ORDER BY id;`, TableName, serverproductparamplayer.TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ServerProductParam
	for rows.Next() {
		c := ServerProductParam{}
		if err = rows.Scan(
			&c.ID,
			&c.DevCoin,
			&c.DevGmp,
			&c.DevItem1,
			&c.DevItem2,
			&c.DevPlatlv01,
			&c.DevPlatlv02,
			&c.DevPlatlv03,
			&c.DevPlatlv04,
			&c.DevPlatlv05,
			&c.DevPlatlv06,
			&c.DevPlatlv07,
			&c.DevRescount01Value,
			&c.DevRescount02Value,
			&c.DevResource01Id,
			&c.DevResource02Id,
			&c.DevSkil,
			&c.DevSpecial,
			&c.DevTime,
			&c.Type,
			&c.UseGmp,
			&c.UseRescount01Value,
			&c.UseRescount02Value,
			&c.UseResource01Id,
			&c.UseResource02Id,
			&c.Open,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, err := r.GetCount(ctx)
	if err != nil {
		return err
	}
	if all > 0 {
		slog.Info("server product params already seeded")
		return nil
	}

	for _, v := range ServerProductParams {
		if err = r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
