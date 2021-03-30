package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TransactionRepositoryPostgres struct {
	db *sqlx.DB
}

func NewTransactionRepositoryPostgres(db *sqlx.DB) *TransactionRepositoryPostgres {
	return &TransactionRepositoryPostgres{db: db}
}

func (t TransactionRepositoryPostgres) BeginTransaction() (*sql.Tx, error) {
	begin, err := t.db.Begin()
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Не удалось создать транзакцию.")
	}
	return begin, nil
}

func (t TransactionRepositoryPostgres) Commit(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		logrus.Error(err)
		return errors.New("Не удалось сделать коммит транзакции")
	}
	return nil
}

func (t TransactionRepositoryPostgres) Rollback(tx *sql.Tx) error {
	if err := tx.Rollback(); err != nil {
		logrus.Error(err)
		return errors.New("Не удалось сделать откат транзакции")
	}
	return nil
}
