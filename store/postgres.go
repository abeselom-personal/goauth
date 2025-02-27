package store

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type PostgresDB struct {
    db *sql.DB
}

func (db *PostgresDB) Connect(connStr string) error {
    var err error
    db.db, err = sql.Open("postgres", connStr)
    return err
}

func (db *PostgresDB) Disconnect() error {
    return db.db.Close()
}
