package database

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

var db *sql.DB

type DBTransaction struct {
	tx      *sql.Tx
	Builder sq.StatementBuilderType
}

func ConnectDB() (err error) {
	db, err = sql.Open("postgres", "host=localhost port=5432 user=dev password=dev123456 dbname=tw sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()

	return err
}

func NewTransaction(ctx context.Context, readOnly bool) (DBtransaction *DBTransaction, err error) {
	DBtransaction = new(DBTransaction)
	DBtransaction.tx, err = db.BeginTx(ctx, &sql.TxOptions{ReadOnly: readOnly})
	if err != nil {
		return DBtransaction, err
	}

	DBtransaction.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)

	return
}

func (t *DBTransaction) Rollback() {
	t.tx.Rollback()
}

func (t *DBTransaction) Commit() (err error) {
	err = t.tx.Commit()
	return
}
