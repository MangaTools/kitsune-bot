package repository

import "database/sql"

type Transaction struct {
	tx *sql.Tx
}

func NewTransaction(tx *sql.Tx) *Transaction {
	return &Transaction{tx: tx}
}

func (t Transaction) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRow(query, args...)
}

func (t Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t Transaction) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.tx.Query(query, args...)
}

func (t Transaction) Begin() (*sql.Tx, error) {
	return t.tx, nil
}

func (t Transaction) Commit() error {
	return t.tx.Commit()
}

func (t Transaction) Rollback() error {
	return t.tx.Rollback()
}
