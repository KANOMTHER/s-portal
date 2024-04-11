package service

import (

	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type ClassService struct {
	db *gorm.DB
}

func NewClassService(db *gorm.DB) *ClassService {
	return &ClassService{
		db: db,
	}
}

func (ps *ClassService) CreateClass(class *model.Class) error {
	if err := ps.db.Create(&class).Error; err != nil {
		return err
	}
	return nil
}