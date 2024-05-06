package dbadapters

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite(dbFilePath string) (conn *sqlx.DB, err error) {
	db, err := sqlx.Open("sqlite3", "file:"+dbFilePath+"?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
