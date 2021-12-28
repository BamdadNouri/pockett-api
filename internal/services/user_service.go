package services

import (
	"gorm.io/gorm"
)

type UserService interface {
	Add()
	Delete()
	Update()
	GetBulk()
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db}
}

func (s *userService) Add() {
}

func (s *userService) Delete() {}

func (s *userService) Update() {}

func (s *userService) GetBulk() {}
