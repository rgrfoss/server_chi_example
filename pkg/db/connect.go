package db

import (
	"database/sql"
	"os"
)

func PostgresConnect() (conn *sql.DB, Err error) {

	conn, err := sql.Open("pgx", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
