package services

import (
	"gorm.io/gorm"
)

type WalletService interface {
	Add()
	Delete()
	Update()
	GetBulk()
}

type walletService struct {
	db *gorm.DB
}

func NewWalletService(db *gorm.DB) *walletService {
	return &walletService{db}
}

func (s *walletService) Add() {
}

func (s *walletService) Delete() {}

func (s *walletService) Update() {}

func (s *walletService) GetBulk() {}
