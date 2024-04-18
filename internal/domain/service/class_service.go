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

type GetClassBySemesterAndYearField struct {
	ID 	uint `example:"1"`
	Section     string `example:"A"`
	CourseID	uint `example:"1"`
	Course	struct {
		ID			uint `example:"1"`
		CourseCode  string `example:"CPE313"`
		CourseName  string `example:"signals and linear systems"`
		Semester	int `example:"2"`
		Year		int `example:"2021"`
	}
}

func (cs *ClassService) GetClassBySemesterAndYear(semester string, year string) ([]GetClassBySemesterAndYearField, error) {
	var classData []GetClassBySemesterAndYearField
	if err := cs.db.
	Model(&model.Class{}).Debug().
	InnerJoins("Course").
	Distinct("classes.ID").
	Select("classes.ID AS ID, classes.section AS Section, classes.course_id AS CourseID").
	Order("classes.ID ASC").
	Where("Course.semester = ? AND Course.year = ?", semester, year).
	Find(&classData).Error; err != nil {
		return nil, err
	}

	return classData, nil
}

func (cs *ClassService) DeleteClassByID(id string) error {
	class := model.Class{}
	if result := cs.db.Delete(&class, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the class")
	}

	return nil
}