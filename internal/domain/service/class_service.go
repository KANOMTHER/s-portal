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

func (cs *ClassService) CreateClass(class *model.Class) error {
	if err := cs.db.Create(&class).Error; err != nil {
		return err
	}
	return nil
}

func (cs *ClassService) GetClassByID(id string) (*model.Class, error) {
	var class *model.Class
	if err := cs.db.Preload("Course").First(&class, id).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (cs *ClassService) GetClassByCourseID(course_id string) ([]model.Class, error) {
	var class []model.Class
	if err := cs.db.Where("course_id = ?", course_id).Preload("Course").Find(&class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (cs *ClassService) GetClassBySemesterAndYear(semester string, year string) ([]uint, error) {
	var class_id []uint
	if err := cs.db.Model(&model.Class{}).Joins("inner join courses on classes.course_id = courses.ID ").Distinct("classes.ID").Order("classes.ID ASC").Where("semester = ? AND year = ?", semester, year).Pluck("classes.ID", &class_id).Error; err != nil {
		return nil, err
	}
	return class_id, nil
}

func (cs *ClassService) DeleteClassByID(id string) error {
	class := model.Class{}
	if result := cs.db.Delete(&class, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the class")
	}

	return nil
}