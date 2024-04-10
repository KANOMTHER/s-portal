package service

import (
	"fmt"

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

func (fs *FacultyService) GetFacultyByID(id string) (*model.Faculty, error) {
	var faculty model.Faculty
	if err := fs.db.First(&faculty, id).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (fs *FacultyService) UpdateFacultyByID(context *gin.Context, id string) error {
	faculty := model.Faculty{}
	if err := fs.db.First(&faculty, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&faculty); err != nil {
		return err
	}

	if err := fs.db.Save(&faculty).Error; err != nil {
		return err
	}

	return nil
}

func (fs *FacultyService) DeleteFacultyByID(id string) error {
	if result := fs.db.Delete(&model.Faculty{}, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the Faculty")
	}
	return nil
}
