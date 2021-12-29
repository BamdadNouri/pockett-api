package repositories

import (
	"fmt"
	"sandbox/pockett-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserEntity struct {
	ID              uint64
	Email           string
	Username        string
	Password        string
	Theme           models.AppTheme
	DefaultWalletID int
	Active          bool
}

type UserRepository interface {
	AddUser(user UserEntity) (*UserEntity, error)
	UpdateUser()
	DeleteUser()
	GetUsers()
	GetByEmail(email string) (*UserEntity, error)
	GetByUsername(username string) (*UserEntity, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &UserRepo{db}
}

func (u *UserRepo) AddUser(user UserEntity) (*UserEntity, error) {
	var res UserEntity

	_, err := u.db.Query(fmt.Sprintf(
		"INSERT INTO users VALUES(0, '%s', '%s', '%s', %d, NULL, %t);",
		user.Email, user.Username, user.Password, user.Theme, user.Active,
	))
	if err != nil {
		return nil, err
	}

	r, err := u.db.Query(
		"SELECT id, email, username, theme, is_active  FROM users ORDER BY id DESC LIMIT 1;",
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		if err := r.Scan(
			&res.ID, &res.Email, &res.Username, &res.Theme, &res.Active,
		); err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (u *UserRepo) UpdateUser() {}

func (u *UserRepo) DeleteUser() {}

func (u *UserRepo) GetUsers() {}

func (u *UserRepo) GetByEmail(email string) (*UserEntity, error) {
	var res UserEntity

	r, err := u.db.Query(
		fmt.Sprintf("SELECT id, email, username, theme, is_active, password  FROM users WHERE email = '%s';", email),
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		if err := r.Scan(
			&res.ID, &res.Email, &res.Username, &res.Theme, &res.Active, &res.Password,
		); err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (u *UserRepo) GetByUsername(username string) (*UserEntity, error) {
	var res UserEntity

	r, err := u.db.Query(
		fmt.Sprintf("SELECT id, email, username, theme, is_active, password  FROM users WHERE username = '%s';", username),
	)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		if err := r.Scan(
			&res.ID, &res.Email, &res.Username, &res.Theme, &res.Active, &res.Password,
		); err != nil {
			return nil, err
		}
	}
	return &res, nil
}
