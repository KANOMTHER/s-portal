package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type FacultyService struct {
	db *gorm.DB
}

func NewFacultyService(db *gorm.DB) *FacultyService {
	return &FacultyService{
		db: db,
	}
}

func (fs *FacultyService) GetAllFaculties() ([]model.Faculty, error) {
	var faculties []model.Faculty
	if err := fs.db.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}

func (fs *FacultyService) CreateFaculty(context *gin.Context, course *model.Faculty) error {
	if err := fs.db.Create(&course).Error; err != nil {
		return err
	}
	return nil
}
