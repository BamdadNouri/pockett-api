package repositories

import (
	"sandbox/pockett-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	AddTransaction(models.TransactionEntity) (*models.TransactionEntity, error)
	UpdateTransaction()
	DeleteTransaction()
	GetTransactions(id, walletID uint64) (*[]models.TransactionEntity, error)
}

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) TransactionRepository {
	return &TransactionRepo{db}
}

func (t *TransactionRepo) AddTransaction(tr models.TransactionEntity) (*models.TransactionEntity, error) {
	var res models.TransactionEntity

	_, err := t.db.Query(
		"INSERT INTO transactions VALUES(0, ?, ?, ?, NULL, ?, ?, false);",
		// "INSERT INTO transactions VALUES(0, %.0f, %d, '%s', 0, %d, false);",
		tr.Amount, tr.TransactionType, tr.Description, tr.Owner, tr.WalletID,
	)
	if err != nil {
		return nil, err
	}
	r, err := t.db.Query(
		"SELECT id, amount, tr_type, description, owner_id, wallet_id, is_deleted FROM transactions ORDER BY id DESC LIMIT 1;",
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		if err := r.Scan(
			&res.ID, &res.Amount, &res.TransactionType, &res.Description, &res.Owner, &res.WalletID, &res.IsDeleted,
		); err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (t *TransactionRepo) UpdateTransaction() {}

func (t *TransactionRepo) DeleteTransaction() {}

func (t *TransactionRepo) GetTransactions(id, walletID uint64) (*[]models.TransactionEntity, error) {
	var respose []models.TransactionEntity

	r, err := t.db.Query(
		"SELECT id, amount, tr_type, description, owner_id, wallet_id, is_deleted FROM transactions WHERE wallet_id = ? AND owner_id = ?;",
		walletID, id,
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		var res models.TransactionEntity
		if err := r.Scan(
			&res.ID, &res.Amount, &res.TransactionType, &res.Description, &res.Owner, &res.WalletID, &res.IsDeleted,
		); err != nil {
			return nil, err
		}
		respose = append(respose, res)
	}
	return &respose, nil
}
