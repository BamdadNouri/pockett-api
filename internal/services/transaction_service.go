package services

import (
	"gorm.io/gorm"
)

type TransactionService interface {
	Add()
	Delete()
	Update()
	GetBulk()
}

type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *transactionService {
	return &transactionService{db}
}

func (s *transactionService) Add() {
	// tag models.TagReq

	// if tag.Action == 0 {
	// 	// spend
	// } else if tag.Action == 1 {
	// 	// add
	// }
	// t := models.TagE{
	// 	Amount:      tag.Amount,
	// 	Action:      tag.Action,
	// 	Description: tag.Description,
	// 	Owner:       "bamdad",
	// }
	// s.db.Create(&t)
	// return t, nil
}

func (s *transactionService) Delete() {}

func (s *transactionService) Update() {}

func (s *transactionService) GetBulk() {}
