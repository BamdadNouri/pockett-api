package repositories

import "github.com/jmoiron/sqlx"

func InitRepositories(db *sqlx.DB) (UserRepository, TransactionRepository, WalletRepository) {
	return NewUserRepo(db), NewTransactionRepo(db), NewWalletRepo(db)
}
