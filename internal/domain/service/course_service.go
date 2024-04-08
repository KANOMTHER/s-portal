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

func (fs *CourseService) GetAllCourses() ([]model.Course, error) {
	var courses []model.Course
	if err := fs.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (fs *CourseService) CreateCourse(context *gin.Context, course *model.Course) error {
	if err := fs.db.Create(&course).Error; err != nil {
		return err
	}
	return nil
}
