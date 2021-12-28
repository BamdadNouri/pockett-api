package services

import (
	"gorm.io/gorm"
)

type TagService interface {
	Add()
	Delete()
	Update()
	GetBulk()
}

type tagService struct {
	db *gorm.DB
}

func NewTagService(db *gorm.DB) *tagService {
	return &tagService{db}
}

func (s *tagService) Add() {
}

func (s *tagService) Delete() {}

func (s *tagService) Update() {}

func (s *tagService) GetBulk() {}
