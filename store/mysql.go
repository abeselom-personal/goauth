package store

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
    db *sql.DB
}

func (db *MySQLDB) Connect(connStr string) error {
    var err error
    db.db, err = sql.Open("mysql", connStr)
    return err
}

func (db *MySQLDB) Disconnect() error {
    return db.db.Close()
}
