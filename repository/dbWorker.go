package repository

import (
	"database/sql"
)

type DbWorker interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Begin() (*sql.Tx, error)
}
