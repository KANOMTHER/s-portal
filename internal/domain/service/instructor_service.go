package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type InstructorService struct {
    db *gorm.DB
}

func NewInstructorService(db *gorm.DB) *InstructorService {
    return &InstructorService{
        db: db,
    }
}

func (is *InstructorService) GetAllInstructors() ([]model.Instructor, error) {
    var Instructors []model.Instructor
    if err := is.db.Find(&Instructors).Error; err != nil {
        return nil, err
    }
    return Instructors, nil
}

func (is *InstructorService) CreateInstructor(instructor *model.Instructor) error {
    if err := is.db.Create(&instructor).Error; err != nil {
        return err
    }
        
    teacherUser := model.GetUserBuilder("teacher")
    director := model.NewUserDirector(teacherUser)
    user := director.Construct(instructor.ID)
    if err := is.db.Create(&user).Error; err != nil {
        return err
    }
    return nil
}

func (is *InstructorService) GetInstructorByID(id string) (*model.Instructor, error) {
    var instructor *model.Instructor
    if err := is.db.
    Preload("Professor.Faculty").
    Preload("Class.Course").
    First(&instructor, id).Error; err != nil {
        return nil, err
    }
    return instructor, nil
}

func (is *InstructorService) UpdateInstructorByID(context *gin.Context, id string) error {
    instructor := model.Instructor{}
    

    if err := is.db.First(&instructor, id).Error; err != nil {
        return err
    }

    if err := context.ShouldBindJSON(&instructor); err != nil {
        return err
    }

    if err := is.db.Save(&instructor).Error; err != nil {
        return err
    }

    return nil
}

func (is *InstructorService) DeleteInstructorByID(id string) error {
    instructor := model.Instructor{}
    if result := is.db.Delete(&instructor, id); result.RowsAffected < 1 {
        return fmt.Errorf("were not able to delete the instructor")
    }

    return nil
}
  