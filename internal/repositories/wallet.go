package repositories

import "github.com/jmoiron/sqlx"

type WalletEntity struct {
	Title  string
	Config string
}

type WalletRepository interface {
	AddWallet()
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

func (t *WalletRepo) AddWallet() {}

func (t *WalletRepo) UpdateWallet() {}

func (t *WalletRepo) DeleteWallet() {}

func (t *WalletRepo) GetWallets() {}
