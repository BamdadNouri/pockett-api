package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UserRepository        UserRepository
	TransactionRepository TransactionRepository
	WalletRepository      WalletRepository
}

var R Repositories

func InitRepositories(db *sqlx.DB) Repositories {
	R = Repositories{
		NewUserRepo(db),
		NewTransactionRepo(db),
		NewWalletRepo(db),
	}
	return R
}
