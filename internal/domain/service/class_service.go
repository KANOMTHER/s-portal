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

func (ps *ClassService) GetClassByID(id string) (*model.Class, error) {
	var class *model.Class
	if err := ps.db.Preload("Course").First(&class, id).Error; err != nil {
		return nil, err
	}
	return class, nil
}
