package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type ProfessorService struct {
	db *gorm.DB
}

func NewProfessorService(db *gorm.DB) *ProfessorService {
	return &ProfessorService{
		db: db,
	}
}

type ProfessorProfile struct {
	ID        uint   `example:"1"`
	FName     string `example:"John"`
	LName     string `example:"Doe"`
	Email     string `example:"Juwan98@example.net"`
	Phone     string `example:"744-512-3072"`
	Position  string `example:"Professor"`
	FacultyID uint   `example:"1"`
	Faculty   struct {
		ID         uint   `example:"1"`
		Major      string `example:"Software"`
		Department string `example:"Engineering"`
	}
}

func (ps *ProfessorService) GetAllProfessors() ([]ProfessorProfile, error) {
	var Professors []ProfessorProfile

	if err := ps.db.Model(&model.Professor{}).Joins("Faculty").Find(&Professors).Error; err != nil {
		return nil, err
	}
	return Professors, nil
}

func (ps *ProfessorService) CreateProfessor(professor *model.Professor) error {
	if err := ps.db.Create(&professor).Error; err != nil {
		return err
	}
	return nil
}

func (ps *ProfessorService) GetProfessorByID(id string) (*ProfessorProfile, error) {
	var professor *ProfessorProfile

	if err := ps.db.Model(&model.Professor{}).Joins("Faculty").First(&professor, id).Error; err != nil {
		return nil, err
	}
	return professor, nil
}

func (ps *ProfessorService) UpdateProfessorByID(context *gin.Context, id string) error {
	professor := model.Professor{}

	if err := ps.db.First(&professor, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&professor); err != nil {
		return err
	}

	if err := ps.db.Save(&professor).Error; err != nil {
		return err
	}

	return nil
}

func (ps *ProfessorService) DeleteProfessorByID(id string) error {
	professor := model.Professor{}

	if result := ps.db.Delete(&professor, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the professor")
	}

	return nil
}
