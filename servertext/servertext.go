package servertext

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ServerText struct {
	ID         int    `json:"-"`
	Identifier string `json:"identifier"`
	Language   string `json:"language"`
	Text       string `json:"text"`
}

var TableName = "serverText"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		identifier TEXT,
		language TEXT,
		text TEXT,
		UNIQUE(identifier,language)
		);
		CREATE INDEX IF NOT EXISTS stext_id_lang ON %s (identifier, language);`,
		TableName, TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *ServerText) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(identifier, language, text) values (?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q, c.Identifier, c.Language, c.Text); err != nil {
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

func (r *Repo) GetAll(ctx context.Context) ([]ServerText, error) {
	q := fmt.Sprintf(`
		SELECT
			id, identifier, language, text
		FROM %s
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ServerText
	for rows.Next() {
		c := ServerText{}
		if err = rows.Scan(&c.ID, &c.Identifier, &c.Language, &c.Text); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (r *Repo) Seed(ctx context.Context) error {
	all, _ := r.GetAll(ctx)
	if len(all) > 0 {
		slog.Info("server texts already seeded")
		return nil
	}

	for _, v := range ServerTexts {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
