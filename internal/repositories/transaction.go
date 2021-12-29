package repositories

import (
	"sandbox/pockett-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	AddTransaction(models.TransactionEntity)
	UpdateTransaction()
	DeleteTransaction()
	GetTransactions()
}

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) TransactionRepository {
	return &TransactionRepo{db}
}

func (t *TransactionRepo) AddTransaction(models.TransactionEntity) {}

func (t *TransactionRepo) UpdateTransaction() {}

func (t *TransactionRepo) DeleteTransaction() {}

func (t *TransactionRepo) GetTransactions() {}
