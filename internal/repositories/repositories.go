package repositories

import "database/sql"

func InitRepositories(db *sql.DB) (UserRepository, TransactionRepository, WalletRepository) {
	return NewUserRepo(db), NewTransactionRepo(db), NewWalletRepo(db)
}
