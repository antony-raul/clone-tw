package config

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (err error) {
	db, err = sql.Open("postgres", "host=localhost port=5432 user=dev password=dev123456 dbname=tw sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()

	return err
}

func NewTransaction(ctx context.Context, readOnly bool) (tx *sql.Tx, err error) {
	tx = new(sql.Tx)
	tx, err = db.BeginTx(ctx, &sql.TxOptions{ReadOnly: readOnly})
	if err != nil {
		return tx, err
	}

	return
}
