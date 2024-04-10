package service

import (

	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type ProgramService struct {
	db *gorm.DB
}

func NewProgramService(db *gorm.DB) *ProgramService {
	return &ProgramService{
		db: db,
	}
}

func (ps *ProgramService) GetAllPrograms() ([]model.Program, error) {
	var Programs []model.Program
	if err := ps.db.Preload("Faculty").Find(&Programs).Error; err != nil {
		return nil, err
	}
	return Programs, nil
}
