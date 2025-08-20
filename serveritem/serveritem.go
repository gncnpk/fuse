package serveritem

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/serverproductparam"
	"log/slog"
)

// ServerItem is a ServerProduct development status
type ServerItem struct {
	ProductID  int
	PlayerID   int
	CreateDate int
	Develop    int
	MbCoin     int
	Open       int // TODO depends on previously developed items, ServerProductParam.DevItem1/2

	Gmp       int
	MaxSecond int
}

// UpdateDevelop sets the develop status for a server item record
func (r *Repo) UpdateDevelop(ctx context.Context, playerID int, productID int, develop int) error {
	q := fmt.Sprintf(`UPDATE %s SET develop = ? WHERE player_id = ? AND product_id = ?;`, TableName)
	_, err := r.db.ExecContext(ctx, q, develop, playerID, productID)
	return err
}

var TableName = "serverItem"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			product_id INTEGER,
			player_id INTEGER,
			create_date INTEGER,
			develop INTEGER,
			mb_coin INTEGER,
			open INTEGER,	
			UNIQUE(product_id, player_id)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ServerItem) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

		q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			product_id,
			player_id,
			create_date,
			develop,
			mb_coin,

			open
		) values (?,?,?,?,?, ?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ProductID,
		c.PlayerID,
		c.CreateDate,
		c.Develop,
		c.MbCoin,

		c.Open,
	); err != nil {
		slog.Error("add failed", "table", TableName, "error", err.Error())
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]ServerItem, error) {
	q := fmt.Sprintf(`
		SELECT
			si.product_id,
			si.player_id,
			si.create_date,
			si.develop,
			si.mb_coin,
			si.open,
			pr.dev_gmp,
			pr.dev_time
		FROM %s si 
		JOIN %s pr on si.product_id = pr.id
		WHERE 
			player_id = ?
		ORDER BY id;`, TableName, serverproductparam.TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ServerItem
	for rows.Next() {
		c := ServerItem{}
		if err = rows.Scan(
			&c.ProductID,
			&c.PlayerID,
			&c.CreateDate,
			&c.Develop,
			&c.MbCoin,
			&c.Open,
			&c.Gmp,
			&c.MaxSecond,
		); err != nil {
			return nil, err
		}

		c.MaxSecond *= 60
		res = append(res, c)
	}

	return res, nil
}

// TODO separate table for nukes
func (r *Repo) GetNukeTime(ctx context.Context, playerID int) (ServerItem, error) {
	q := fmt.Sprintf(`
		SELECT
			si.product_id,
			si.player_id,
			si.create_date,
			si.develop,
			si.mb_coin,
			si.open,
			pr.dev_gmp,
			pr.dev_time
		FROM %s si 
		JOIN %s pr on si.product_id = pr.id
		WHERE 
			player_id = ? and si.create_date > 0;`, TableName, serverproductparam.TableName)
	rows := r.db.QueryRowContext(ctx, q, playerID)
	c := ServerItem{}
	if err := rows.Scan(
		&c.ProductID,
		&c.PlayerID,
		&c.CreateDate,
		&c.Develop,
		&c.MbCoin,
		&c.Open,
		&c.Gmp,
		&c.MaxSecond,
	); err != nil {
		return ServerItem{}, err
	}

	c.MaxSecond *= 60

	return c, nil
}
