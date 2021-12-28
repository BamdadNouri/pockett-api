package modules

import (
	"sandbox/pockett-api/internal/models"
	"sandbox/pockett-api/internal/repositories"
)

type Transaction interface {
	Create(transaction models.TransactionCreateReq) Transaction
	Update(transaction models.TransactionUpdateReq) Transaction
	Find(id uint64) Transaction
	Bulk(page, size int) []Transaction
	SoftDelete(id uint64) Transaction
	Result() (Transaction, error)
	Error() error
	ToRes() models.TransactionRes
}

type transaction struct {
	repository repositories.TransactionRepo

	id              uint64
	ownerID         uint64
	transactionType models.TransactionType
	amount          float64
	description     string
	tagIDs          []uint64
	isDeleted       bool

	err error
}

func NewTransaction(ownerID uint64) Transaction {
	return &transaction{
		ownerID: ownerID,
	}
}

func (t *transaction) Create(transaction models.TransactionCreateReq) Transaction {
	return t
}
func (t *transaction) Update(transaction models.TransactionUpdateReq) Transaction {
	return t
}
func (t *transaction) Find(id uint64) Transaction {
	return t
}
func (t *transaction) Bulk(page, size int) []Transaction {
	return []Transaction{t}
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
		TagIDs:          t.tagIDs,
	}
}
