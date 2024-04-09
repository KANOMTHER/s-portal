package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type CourseService struct {
	db *gorm.DB
}

func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{
		db: db,
	}
}

func (cs *CourseService) GetAllCourses() ([]model.Course, error) {
	var courses []model.Course
	if err := cs.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (cs *CourseService) CreateCourse(context *gin.Context, course *model.Course) error {
	if err := cs.db.Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CourseService) FindCourseByID(context *gin.Context, id string) (*model.Course, error) {
	var course *model.Course
	if err := cs.db.First(&course, id).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (cs *CourseService) UpdateCourseByID(context *gin.Context, id string) error {
	course := model.Course{}
	if err := cs.db.First(&course, id).Error; err != nil {
		return err
	}

	if err := context.ShouldBindJSON(&course); err != nil {
		return err
	}

	if err := cs.db.Save(&course).Error; err != nil {
		return err
	}

	return nil
}
