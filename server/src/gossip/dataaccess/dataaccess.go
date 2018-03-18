package dataaccess

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

func NewConnection(url string) (conn Connection, err error) {
	if conn.db, err = sql.Open("postgres", url); err == nil {
		err = conn.db.Ping()
	}
	return
}
