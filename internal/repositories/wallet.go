package repositories

import (
	"sandbox/pockett-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type WalletEntity struct {
	Title  string
	Config string
}

type WalletRepository interface {
	AddWallet(wallet models.WalletEntity) (*models.WalletEntity, error)
	UpdateWallet()
	DeleteWallet()
	GetWallets()
}

type WalletRepo struct {
	db *sqlx.DB
}

func NewWalletRepo(db *sqlx.DB) WalletRepository {
	return &WalletRepo{db}
}

func (t *WalletRepo) AddWallet(wallet models.WalletEntity) (*models.WalletEntity, error) {
	var res models.WalletEntity

	_, err := t.db.Query(
		"INSERT INTO wallets VALUES(0, ?, ?, false);",
		wallet.Title, wallet.OwnerID,
	)
	if err != nil {
		return nil, err
	}
	r, err := t.db.Query(
		"SELECT id, title, owner_id FROM wallets ORDER BY id DESC LIMIT 1;",
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		if err := r.Scan(
			&res.ID, &res.Title, &res.OwnerID,
		); err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (t *WalletRepo) UpdateWallet() {}

func (t *WalletRepo) DeleteWallet() {}

func (t *WalletRepo) GetWallets() {}
