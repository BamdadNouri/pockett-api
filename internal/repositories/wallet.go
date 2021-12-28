package repositories

import "database/sql"

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
	db *sql.DB
}

func NewWalletRepo(db *sql.DB) WalletRepository {
	return &WalletRepo{db}
}

func (t *WalletRepo) AddWallet() {}

func (t *WalletRepo) UpdateWallet() {}

func (t *WalletRepo) DeleteWallet() {}

func (t *WalletRepo) GetWallets() {}
