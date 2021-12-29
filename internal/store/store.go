package store

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Init() Store
	RunMigration() Store
	Result() (*sqlx.DB, error)
}

type store struct {
	db      *sqlx.DB
	dsn     string
	migrate bool

	err error
}

func NewStore(dsn string, migrate bool) Store {
	return &store{
		dsn:     dsn,
		migrate: migrate,
	}
}

func (s *store) Init() Store {
	fmt.Println("init store", s.dsn)
	db, err := sqlx.Open("mysql", s.dsn)
	// db, err := sql.Open("mysql", s.dsn)
	if err != nil {
		s.err = fmt.Errorf("error in initializing database", err)
		return s
	}
	r, err := db.Query("show tables;")
	if err != nil {
		s.err = fmt.Errorf("error in getting tables", err)
		return s
	}
	tables := []string{}
	defer r.Close()
	for r.Next() {
		var t string
		if err := r.Scan(
			&t,
		); err != nil {
			s.err = fmt.Errorf("error in scanning table names", err)
			return s
		}
		tables = append(tables, t)
	}
	if len(tables) > 0 {
		s.migrate = false
	}
	s.db = db
	return s
}

func (s *store) RunMigration() Store {
	if s.migrate {
		if s.err != nil {
			return s
		}
		tables := []string{
			CreateUsersTable,
			CreateTagsTable,
			CreateWalletsTable,
			CreateTransactionsTable,
			CreateUserRecord,
			CreateWalletRecord,
			CreateTransactionRecord}
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

func (s *store) Result() (*sqlx.DB, error) {
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
