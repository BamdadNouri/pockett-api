package repositories

import (
	"database/sql"
	"sandbox/pockett-api/internal/models"
)

type TransactionRepository interface {
	AddTransaction(models.TransactionEntity)
	UpdateTransaction()
	DeleteTransaction()
	GetTransactions()
}

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepository {
	return &TransactionRepo{db}
}

func (t *TransactionRepo) AddTransaction(models.TransactionEntity) {}

func (t *TransactionRepo) UpdateTransaction() {}

func (t *TransactionRepo) DeleteTransaction() {}

func (t *TransactionRepo) GetTransactions() {}
