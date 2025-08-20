package tppmessage

import (
	"context"
	"database/sql"
	"fmt"
)

type URLListEntry struct {
	Type    string `json:"type"`
	Url     string `json:"url"`
	Version int    `json:"version"`
}

type CMDGetURLListResponse struct {
	CryptoType string         `json:"crypto_type"`
	Flowid     interface{}    `json:"flowid"`
	Msgid      string         `json:"msgid"`
	Result     string         `json:"result"`
	Rqid       int            `json:"rqid"`
	UrlList    []URLListEntry `json:"url_list"`
	UrlNum     int            `json:"url_num"`
	Xuid       interface{}    `json:"xuid"`
}

type CMDGetURLListRequest struct {
	Lang   string `json:"lang"`
	Msgid  string `json:"msgid"`
	Region string `json:"region"`
	Rqid   int    `json:"rqid"`
}

type URLListEntryRepo struct {
	db *sql.DB
}

func (u *URLListEntryRepo) Clear(ctx context.Context) error {
	query := fmt.Sprintf(`DELETE FROM %s`, u.TableName())
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *URLListEntryRepo) Insert(ctx context.Context, entry URLListEntry) error {
	query := fmt.Sprintf(`insert into %s (urlType, url, version) values(?,?,?)`, u.TableName())
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, entry.Type, entry.Url, entry.Version)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *URLListEntryRepo) GetAll(ctx context.Context) ([]URLListEntry, error) {
	query := fmt.Sprintf(`select urlType, url, version from %s`, u.TableName())
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot get all url entries: %w", err)
	}
	defer rows.Close()

	res := make([]URLListEntry, 0)
	for rows.Next() {
		var e URLListEntry
		if err = rows.Scan(&e.Type, &e.Url, &e.Version); err != nil {
			return nil, fmt.Errorf("cannot scan urlEntry: %w", err)
		}
		res = append(res, e)
	}

	return res, nil
}

func (u *URLListEntryRepo) TableName() string {
	return "urlListEntry"
}

func (u *URLListEntryRepo) WithDB(DB *sql.DB) *URLListEntryRepo {
	u.db = DB
	return u
}

func (u *URLListEntryRepo) CreateSchema() error {
	schema := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	urlType TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL,
	version INTEGER
	);`, u.TableName())

	_, err := u.db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
