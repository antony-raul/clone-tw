package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev password=dev123456 dbname=tw sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
