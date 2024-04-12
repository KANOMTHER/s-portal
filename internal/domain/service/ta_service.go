package service

import (
	//"fmt"

	//"github.com/gin-gonic/gin"
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

func (ts *TAService) GetSimpleText() (cout *string) {
	text := "parnnoi"
	return &text
}

func (ts *TAService) CreateTA(context *gin.Context) (err error) {
	data := model.TA{}

    if err := context.ShouldBindJSON(&data); err != nil {
        return err
    }

	var existingTA model.TA
	if err := ts.db.Where(&data).First(&existingTA).Error; err != nil {
		if(err == gorm.ErrRecordNotFound) {
			factory := model.StaffFactory{}
			ta := factory.CreateAssistance(data.ClassID, data.StudentID)
			if err := ts.db.Create(&ta).Error; err != nil {
				return err
			}
			return nil
		}
    }

	return fmt.Errorf("this student is already exists in the class")
}

func (ts *TAService) UpdateTA(context *gin.Context) (err error) {
	// id (for find), cid & sid (for edit)
	data := model.TA{}

	id := context.Param("id")

	if err := ts.db.First(&data, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&data); err != nil {
        return err
    }

	if err := ts.db.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (ts *TAService) GetTA() model.TA {
	student := model.TA{StudentID: 1}
	ts.db.Find(&student)

	return student
}