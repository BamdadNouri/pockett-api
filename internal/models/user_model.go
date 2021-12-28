package models

type AppTheme int

const (
	LightTheme AppTheme = 0
	DarkTheme  AppTheme = 1
)

type UserEntity struct {
	ID            uint64
	Email         string
	Username      string
	Password      string
	Theme         AppTheme
	DefaultWallet uint64
	Active        bool
}

type UserCreateReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	ID       string `json:"id"` // username or email
	Password string `json:"password"`
}

type UserUpdateReq struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRes struct {
	ID            uint64   `json:"id"`
	Email         string   `json:"email"`
	Username      string   `json:"username"`
	Theme         AppTheme `json:"theme"`
	DefaultWallet uint64   `json:"defaultWallet"`
	Active        bool     `json:"active"`
}
