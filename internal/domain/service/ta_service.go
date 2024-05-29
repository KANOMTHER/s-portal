package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type TAService struct {
	db *gorm.DB
}

func NewTAService(db *gorm.DB) *TAService {
	return &TAService{
		db: db,
	}
}

func (ts *TAService) GetTA(context *gin.Context) ([]model.TA, error) {
	ta := []model.TA{}

	if err := ts.db.Model(model.TA{}).Find(&ta).Error; err != nil {
		return nil, err
	}

	return ta, nil
}

func (ts *TAService) GetTAByClassID(context *gin.Context) ([]model.TA, error) {
	ta := []model.TA{}

	class_id := context.Param("id")

	if err := ts.db.Model(model.TA{}).Where("class_id = ?", class_id).Find(&ta).Error; err != nil {
		return nil, err
	}

	return ta, nil
}

func (ts *TAService) CreateTA(context *gin.Context) error {
	ta := model.TA{}

	if err := context.ShouldBindJSON(&ta); err != nil {
		return err
	}

	var existingTA model.TA
	if err := ts.db.Where(&ta).First(&existingTA).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			factory := model.StaffFactory{}
			ta := factory.CreateAssistance(ta.ClassID, ta.StudentID)
			if err := ts.db.Create(&ta).Error; err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("this student is already exists in the class")
}

func (ts *TAService) UpdateTA(context *gin.Context) error {
	// id (for find), cid & sid (for edit)
	ta := model.TA{}

	id := context.Param("id")

	if err := ts.db.First(&ta, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&ta); err != nil {
		return err
	}

	invalidInput := model.TA{}

	err := ts.db.Model(model.TA{}).Where("class_id = ? AND student_id = ?", ta.ClassID, ta.StudentID).First(&invalidInput).Error

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("your student_id and class_id is already exists, please try others")
	}

	if err := ts.db.Save(&ta).Error; err != nil {
		return err
	}

	return nil
}

func (ts *TAService) DeleteTA(context *gin.Context) error {
	// id (for find), cid & sid (for edit)
	ta := model.TA{}

	id := context.Param("id")

	if result := ts.db.Delete(&ta, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete this ta")
	}

	return nil
}
