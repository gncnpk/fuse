package informationmessage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type InformationMessage struct {
	ID         int
	InfoID     int
	Date       int
	Important  bool
	MesBody    string
	MesSubject string
	Language   string
	Region     string
}

var TableName = "informationMessage"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		info_id INTEGER,
		date INTEGER,
		important INTEGER,
		body TEXT,
		subject TEXT,
		language TEXT,
		region TEXT
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) Add(ctx context.Context, c *InformationMessage) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`insert into %s(
		id,
		info_id,
		date,
		important,
		body,
		subject,
		language,
		region
		) values (?,?,?,?,?, ?,?,?);`, TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.ID,
		c.InfoID,
		c.Date,
		c.Important,
		c.MesBody,
		c.MesSubject,
		c.Language,
		c.Region,
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

func (r *Repo) GetByRegionLang(ctx context.Context, region string, language string) ([]InformationMessage, error) {
	q := fmt.Sprintf(`
		SELECT
			id,
			info_id,
			date,
			important,
			body,

			subject,
			language,
			region
		FROM %s
		WHERE 
			region = ?
			AND language = ?
		ORDER BY id;`, TableName)
	rows, err := r.db.QueryContext(ctx, q, region, language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []InformationMessage
	for rows.Next() {
		c := InformationMessage{}
		if err = rows.Scan(
			&c.ID,
			&c.InfoID,
			&c.Date,
			&c.Important,
			&c.MesBody,
			&c.MesSubject,
			&c.Language,
			&c.Region,
		); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
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

func (r *Repo) Seed(ctx context.Context) error {
	count, _ := r.GetCount(ctx)
	if count > 0 {
		slog.Info("information messages already seeded")
		return nil
	}

	for _, v := range InformationMessages {
		if err := r.Add(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
