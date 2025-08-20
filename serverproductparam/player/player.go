package serverproductparamplayer

import (
	"context"
	"database/sql"
	"fmt"
)

type ServerProductParamPlayer struct {
	PlayerID  int
	ProductID int
	Open      int
}

var TableName = "serverProductParamPlayer"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		player_id INTEGER,
		product_id INTEGER,
		open INTEGER,
		UNIQUE(player_id, product_id)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ServerProductParamPlayer) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		player_id,
		product_id, 
		open
		) values (?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.ProductID,
		c.Open,
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

func (r *Repo) GetByPlayerID(ctx context.Context, playerID int) ([]ServerProductParamPlayer, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id, product_id, open
		FROM %s
		WHERE
			player_id = ?
		ORDER BY product_id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ServerProductParamPlayer
	for rows.Next() {
		c := ServerProductParamPlayer{}
		if err = rows.Scan(
			&c.PlayerID,
			&c.ProductID,
			&c.Open,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}
