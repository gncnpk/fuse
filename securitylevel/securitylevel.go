package securitylevel

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type SecurityLevel struct {
	PlayerID int
	Slot01   int
	Slot02   int
	Slot03   int
	Slot04   int
	Slot05   int
	Slot06   int
	Slot07   int
	Slot08   int
	Slot09   int
	Slot10   int
	Slot11   int
	Slot12   int
	Slot13   int
	Slot14   int
	Slot15   int
	Slot16   int
	Slot17   int
	Slot18   int
}

func (f *SecurityLevel) FromArray(in []int) error {
	if len(in) != 18 {
		return fmt.Errorf("invalid input length: %d, want 18", len(in))
	}

	f.Slot01 = in[0]
	f.Slot02 = in[1]
	f.Slot03 = in[2]
	f.Slot04 = in[3]
	f.Slot05 = in[4]
	f.Slot06 = in[5]
	f.Slot07 = in[6]
	f.Slot08 = in[7]
	f.Slot09 = in[8]
	f.Slot10 = in[9]
	f.Slot11 = in[10]
	f.Slot12 = in[11]
	f.Slot13 = in[12]
	f.Slot14 = in[13]
	f.Slot15 = in[14]
	f.Slot16 = in[15]
	f.Slot17 = in[16]
	f.Slot18 = in[17]

	return nil
}

func (f *SecurityLevel) ToArray() []int {
	return []int{
		f.Slot01,
		f.Slot02,
		f.Slot03,
		f.Slot04,
		f.Slot05,
		f.Slot06,
		f.Slot07,
		f.Slot08,
		f.Slot09,
		f.Slot10,
		f.Slot11,
		f.Slot12,
		f.Slot13,
		f.Slot14,
		f.Slot15,
		f.Slot16,
		f.Slot17,
		f.Slot18,
	}
}

var TableName = "securityLevel"

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
		slot_08 INTEGER,
		slot_09 INTEGER,
		slot_10 INTEGER,
		slot_11 INTEGER,
		slot_12 INTEGER,
		slot_13 INTEGER,
		slot_14 INTEGER,
		slot_15 INTEGER,
		slot_16 INTEGER,
		slot_17 INTEGER,
		slot_18 INTEGER
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *SecurityLevel) error {
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
			slot_08,
			slot_09,
			slot_10,
			slot_11,
			slot_12,
			slot_13,
			slot_14,
			slot_15,
			slot_16,
			slot_17,
			slot_18
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?);`,
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
		c.Slot09,
		c.Slot10,
		c.Slot11,
		c.Slot12,
		c.Slot13,
		c.Slot14,
		c.Slot15,
		c.Slot16,
		c.Slot17,
		c.Slot18,
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

func (r *Repo) Update(ctx context.Context, c *SecurityLevel) error {
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
			slot_08 = ?,
			slot_09 = ?,
			slot_10 = ?,
			slot_11 = ?,
			slot_12 = ?,
			slot_13 = ?,
			slot_14 = ?,
			slot_15 = ?,
			slot_16 = ?,
			slot_17 = ?,
			slot_18 = ?
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
		c.Slot09,
		c.Slot10,
		c.Slot11,
		c.Slot12,
		c.Slot13,
		c.Slot14,
		c.Slot15,
		c.Slot16,
		c.Slot17,
		c.Slot18,
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

func (r *Repo) Get(ctx context.Context, player_id int) (SecurityLevel, error) {
	q := fmt.Sprintf(`
		SELECT
			player_id,
			slot_01,
			slot_02,
			slot_03,
			slot_04,
			slot_05,
			slot_06,
			slot_07,
			slot_08,
			slot_09,
			slot_10,
			slot_11,
			slot_12,
			slot_13,
			slot_14,
			slot_15,
			slot_16,
			slot_17,
			slot_18
		FROM %s 
		WHERE 
			player_id = ?;`, TableName)
	rows := r.db.QueryRowContext(ctx, q, player_id)

	c := SecurityLevel{}
	if err := rows.Scan(
		&c.PlayerID,
		&c.Slot01,
		&c.Slot02,
		&c.Slot03,
		&c.Slot04,
		&c.Slot05,
		&c.Slot06,
		&c.Slot07,
		&c.Slot08,
		&c.Slot09,
		&c.Slot10,
		&c.Slot11,
		&c.Slot12,
		&c.Slot13,
		&c.Slot14,
		&c.Slot15,
		&c.Slot16,
		&c.Slot17,
		&c.Slot18,
	); err != nil {
		return SecurityLevel{}, err
	}

	return c, nil
}
