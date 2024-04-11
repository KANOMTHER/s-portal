package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

func (ps *ProgramService) CreateProgram(program *model.Program) error {
	if err := ps.db.Create(&program).Error; err != nil {
		return err
	}
	return nil
}

func (ps *ProgramService) GetProgramByID(id string) (*model.Program, error) {
	var program *model.Program
	if err := ps.db.Preload("Faculty").First(&program, id).Error; err != nil {
		return nil, err
	}
	return program, nil
}

func (ps *ProgramService) UpdateProgramByID(context *gin.Context, id string) error {
	program := model.Program{}
	

	if err := ps.db.First(&program, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&program); err != nil {
		return err
	}

	if err := ps.db.Save(&program).Error; err != nil {
		return err
	}

	return nil
}

func (ps *ProgramService) DeleteProgramByID(id string) error {
	program := model.Program{}
	if result := ps.db.Delete(&program, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the program")
	}

	return nil
}