package service

import (
	"fmt"

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

func (ps *ClassService) GetClassByCourseID(id string) ([]model.Class, error) {
	var class []model.Class
	if err := ps.db.Where("course_id = ?", id).Preload("Course").Find(&class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (ps *ClassService) DeleteClassByID(id string) error {
	class := model.Class{}
	if result := ps.db.Delete(&class, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the class")
	}

	return nil
}