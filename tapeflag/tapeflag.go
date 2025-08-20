package tapeflag

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type TapeFlag struct {
	PlayerID int
	Slot01   int
	Slot02   int
	Slot03   int
	Slot04   int
	Slot05   int
	Slot06   int
	Slot07   int
	Slot08   int
}

func (f *TapeFlag) FromArray(in []int) error {
	if len(in) != 8 {
		return fmt.Errorf("invalid input length: %d, want 8", len(in))
	}

	f.Slot01 = in[0]
	f.Slot02 = in[1]
	f.Slot03 = in[2]
	f.Slot04 = in[3]
	f.Slot05 = in[4]
	f.Slot06 = in[5]
	f.Slot07 = in[6]
	f.Slot08 = in[7]
	return nil
}

var TableName = "tapeFlag"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		player_id INTEGER PRIMARY KEY,
		slot_01 INTEGER,
		slot_02 INTEGER,
		slot_03 INTEGER,
		slot_04 INTEGER,
		slot_05 INTEGER,
		slot_06 INTEGER,
		slot_07 INTEGER,
		slot_08 INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *TapeFlag) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			slot_01,
			slot_02,
			slot_03,
			slot_04,
			slot_05,
			slot_06,
			slot_07,
			slot_08
		) values (?,?,?,?,?, ?,?,?,?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.Slot01,
		c.Slot02,
		c.Slot03,
		c.Slot04,
		c.Slot05,
		c.Slot06,
		c.Slot07,
		c.Slot08,
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

func (r *Repo) Update(ctx context.Context, c *TapeFlag) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`UPDATE %s SET 
			slot_01 = ?,
			slot_02 = ?,
			slot_03 = ?,
			slot_04 = ?,
			slot_05 = ?,
			slot_06 = ?,
			slot_07 = ?,
			slot_08 = ?
		WHERE player_id = ?;`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.Slot01,
		c.Slot02,
		c.Slot03,
		c.Slot04,
		c.Slot05,
		c.Slot06,
		c.Slot07,
		c.Slot08,
		c.PlayerID,
	); err != nil {
		slog.Error("update fail", "error", err.Error(), "table", TableName, "playerID", c.PlayerID)
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
