package modules

import (
	"fmt"
	"sandbox/pockett-api/internal/errs"
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/repositories"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CheckExistance(email, username string) User
	Register(user models.UserCreateReq) User
	Login(user models.UserLogin) User
	Find(id uint64) User
	Result() (User, error)
	Token() string
	Error() error
	ToRes() models.UserRes
}

type claims struct {
	ID string `json:"user_id"`
	jwt.StandardClaims
}

type user struct {
	id            uint64
	email         string
	username      string
	password      string
	theme         models.AppTheme
	defaultWallet uint64
	active        bool
	token         string

	userRepo repositories.UserRepository

	err error
}

func NewUser(userRepo repositories.UserRepository) User {
	return &user{userRepo: userRepo}
}

func (u *user) CheckExistance(email, username string) User {
	checkEmail, err := u.userRepo.CheckExistanceByEmail(email)
	if err != nil {
		u.err = err
		return u
	}
	if checkEmail {
		u.err = errs.ErrEmailTaken
		return u
	}
	checkUsername, err := u.userRepo.CheckExistanceByUsername(username)
	if err != nil {
		u.err = err
		return u
	}
	if checkUsername {
		u.err = errs.ErrUsernameTaken
		return u
	}
	return u
}

func (u *user) Register(user models.UserCreateReq) User {
	if u.err != nil {
		return u
	}
	pw, err := u.hashPassword(user.Password)
	if err != nil {
		u.err = err
		return u
	}
	newUser, err := u.userRepo.AddUser(models.UserEntity{
		Email:    user.Email,
		Username: user.Username,
		Password: pw,
		Theme:    models.DarkTheme,
		Active:   true,
	})
	if err != nil {
		u.err = err
		return u
	}
	u.id = newUser.ID
	u.email = newUser.Email
	u.username = newUser.Username
	u.password = newUser.Password
	u.theme = newUser.Theme
	u.active = newUser.Active
	u.generateToken()
	w, err := NewWallet(u.id).AddWallet(models.WalletCreateReq{Title: "Wallet"}).Result()
	if err != nil {
		u.err = err
		return u
	}
	u.setDefaultWallet(w.ToRes().ID)

	return u
}

func (u *user) Login(user models.UserLogin) User {
	if u.err != nil {
		return u
	}
	ue := &models.UserEntity{}
	var err error
	if strings.Contains(user.ID, "@") {
		ue, err = u.userRepo.GetByEmail(user.ID)
	} else {
		ue, err = u.userRepo.GetByUsername(user.ID)
	}
	if err != nil {
		u.err = err
		return u
	}
	if ue.Password == "" {
		u.err = errs.ErrNotFound
		return u
	}
	if !u.checkPasswordHash(user.Password, ue.Password) {
		u.err = errs.ErrAccessDenied
		return u
	}
	u.id = ue.ID
	u.email = ue.Email
	u.username = ue.Username
	u.password = ue.Password
	u.theme = ue.Theme
	u.active = ue.Active
	u.defaultWallet = ue.DefaultWallet
	u.generateToken()

	return u
}

func (u *user) Find(id uint64) User {
	return u
}

func (u *user) AddDefaultWallet(id uint64) User {
	if u.err != nil {
		return u
	}
	wallet, _ := NewWallet(u.id).
		AddWallet(models.WalletCreateReq{Title: "Wallet"}).
		Result()

	u.defaultWallet = wallet.ToRes().ID
	return u
}

func (u *user) Result() (User, error) {
	return u, u.err
}

func (u *user) Token() string {
	return u.token
}

func (u *user) Error() error {
	return u.err
}

func (u *user) ToRes() models.UserRes {
	return models.UserRes{
		ID:            u.id,
		Email:         u.email,
		Username:      u.username,
		Theme:         u.theme,
		DefaultWallet: u.defaultWallet,
		Active:        u.active,
	}
}

func (u *user) setDefaultWallet(id uint64) User {
	err := u.userRepo.SetDefaultWallet(u.id, id)
	if err != nil {
		u.err = err
		return u
	}
	u.defaultWallet = id
	return u
}

func (u *user) generateToken() User {
	token, _ := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims{ID: fmt.Sprintf("%d", u.id)}).
		SignedString([]byte("SECRET"))

	u.token = token
	return u
}

func (u *user) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u *user) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
