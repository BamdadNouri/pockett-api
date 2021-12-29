package modules

import (
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/repositories"
)

type Wallet interface {
	AddWallet(wallet models.WalletCreateReq) Wallet
	Result() (Wallet, error)
	Error() error
	ToRes() models.WalletRes
}

type wallet struct {
	repository repositories.WalletRepository

	id      uint64
	title   string
	color   string
	ownerID uint64

	err error
}

func NewWallet(ownerID uint64) Wallet {
	return &wallet{
		ownerID:    ownerID,
		repository: repositories.R.WalletRepository,
	}
}

func (w *wallet) AddWallet(wallet models.WalletCreateReq) Wallet {
	newWallet, err := w.repository.AddWallet(models.WalletEntity{
		Title:   wallet.Title,
		OwnerID: w.ownerID,
	})
	if err != nil {
		w.err = err
		return w
	}
	w.id = newWallet.ID
	w.title = newWallet.Title

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
