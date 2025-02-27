package store

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
    db *sql.DB
}

func (db *SQLiteDB) Connect(dataSourceName string) error {
    var err error
    db.db, err = sql.Open("sqlite3", dataSourceName)
    return err
}

func (db *SQLiteDB) Disconnect() error {
    return db.db.Close()
}
