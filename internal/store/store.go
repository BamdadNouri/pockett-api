package store

import (
	"database/sql"
	"fmt"
)

type Store interface {
	Init() Store
	RunMigration(migrate bool) Store
	Result() (*sql.DB, error)
}

type store struct {
	db  *sql.DB
	dsn string

	err error
}

func NewStore(dsn string) Store {
	return &store{
		dsn: dsn,
	}
}

func (s *store) Init() Store {
	db, err := sql.Open("mysql", s.dsn)
	if err != nil {
		s.err = fmt.Errorf("error in initializing database", err)
		return s
	}
	s.db = db
	return s
}

func (s *store) RunMigration(migrate bool) Store {
	if migrate {
		if s.err != nil {
			return s
		}
		tables := []string{CreateUsersTable, CreateTagsTable, CreateTransactionsTable, CreateWalletsTable, CreateUserRecord, CreateWalletRecord, CreateTransactionRecord}
		for _, q := range tables {
			_, err := s.db.Query(q)
			if err != nil {
				s.err = err
				fmt.Println("error in migration", err)
				return s
			}
		}
		fmt.Println("done migrating")
	}
	return s
}

func (s *store) Result() (*sql.DB, error) {
	return s.db, s.err
}

// import (
// 	"fmt"
// 	"sandbox/pockett-api/config"
// 	"sandbox/pockett-api/internal/repositories"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // Database is the database struct
// type Database struct {
// }

// // NewDatabase returns a new DB
// func NewDatabase(config *config.Config) (*gorm.DB, error) {
// 	db, err := gorm.Open(postgres.Open(
// 		fmt.Sprintf(
// 			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
// 			config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port, config.Database.SSLMode,
// 		),
// 	), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// err = db.AutoMigrate(&repositories.TransactionEntity{})
// 	return db, err
// }
