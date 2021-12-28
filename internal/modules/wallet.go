package modules

import "sandbox/pockett-api/internal/models"

type Wallet interface {
	AddWallet(wallet models.WalletCreateReq) Wallet
	Result() (Wallet, error)
	Error() error
	ToRes() models.WalletRes
}

type wallet struct {
	id      uint64
	title   string
	color   string
	ownerID uint64

	err error
}

func NewWallet(ownerID uint64) Wallet {
	return &wallet{
		ownerID: ownerID,
	}
}

func (w *wallet) AddWallet(wallet models.WalletCreateReq) Wallet {
	return w
}

func (w *wallet) Result() (Wallet, error) {
	return w, w.err
}

func (w *wallet) Error() error {
	return w.err
}

func (w *wallet) ToRes() models.WalletRes {
	return models.WalletRes{
		ID:    w.id,
		Title: w.title,
	}
}
