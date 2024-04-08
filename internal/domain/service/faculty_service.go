package service

import (
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
