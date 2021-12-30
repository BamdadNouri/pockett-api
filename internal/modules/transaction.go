package modules

import (
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/repositories"
)

type Transaction interface {
	Create(transaction models.TransactionCreateReq) Transaction
	Update(transaction models.TransactionUpdateReq) Transaction
	Find(id uint64, walletID uint64) Transaction
	Bulk(walletID uint64, page, size int) []models.TransactionRes
	SoftDelete(id uint64) Transaction
	Result() (Transaction, error)
	Error() error
	ToRes() models.TransactionRes
	GetBalance() float64
}

type transaction struct {
	repository repositories.TransactionRepository

	id              uint64
	ownerID         uint64
	transactionType models.TransactionType
	amount          float64
	description     string
	walletID        uint64
	tagIDs          []uint64
	isDeleted       bool

	err error
}

func NewTransaction(ownerID uint64, repository repositories.TransactionRepository) Transaction {
	return &transaction{
		ownerID:    ownerID,
		repository: repository,
	}
}

func (t *transaction) Create(transaction models.TransactionCreateReq) Transaction {
	newTR, err := t.repository.AddTransaction(models.TransactionEntity{
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		Description:     transaction.Description,
		Owner:           t.ownerID,
		WalletID:        transaction.WalletID,
		IsDeleted:       false,
	})
	if err != nil {
		t.err = err
		return t
	}
	t.id = newTR.ID
	t.ownerID = newTR.Owner
	t.transactionType = newTR.TransactionType
	t.amount = newTR.Amount
	t.description = newTR.Description
	t.isDeleted = newTR.IsDeleted
	t.walletID = newTR.WalletID

	return t
}
func (t *transaction) Update(transaction models.TransactionUpdateReq) Transaction {
	return t
}
func (t *transaction) Find(id uint64, walletID uint64) Transaction {
	return t
}
func (t *transaction) Bulk(walletID uint64, page, size int) []models.TransactionRes {
	var res []models.TransactionRes
	t.walletID = walletID
	transactions, err := t.repository.GetTransactions(t.ownerID, walletID)
	if err != nil || len(*transactions) == 0 {
		t.err = err
		return []models.TransactionRes{}
	}
	for _, tr := range *transactions {
		res = append(res, models.TransactionRes{
			ID:              tr.ID,
			Amount:          tr.Amount,
			TransactionType: tr.TransactionType,
			Description:     tr.Description,
			WalletID:        tr.WalletID,
		})
	}
	return res
}
func (t *transaction) SoftDelete(id uint64) Transaction {
	return t
}
func (t *transaction) Result() (Transaction, error) {
	return t, t.err
}
func (t *transaction) Error() error {
	return t.err
}
func (t *transaction) ToRes() models.TransactionRes {
	return models.TransactionRes{
		ID:              t.id,
		Amount:          t.amount,
		TransactionType: t.transactionType,
		Description:     t.description,
		WalletID:        t.walletID,
		// TagIDs:          t.tagIDs,
	}
}

func (t *transaction) GetBalance() float64 {
	balanceDetails, err := t.repository.GetBalanceDetails(t.ownerID, t.walletID)
	if err != nil {
		t.err = err
		return 0
	}
	return balanceDetails[models.Earn] - balanceDetails[models.Spend]
}
