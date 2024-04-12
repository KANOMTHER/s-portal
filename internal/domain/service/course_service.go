package service

import (
	"fmt"

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

func (cs *CourseService) GetAllDistinctSemester() ([]int, error) {
	var courses []model.Course
	var semesters []int
	if err := cs.db.Distinct("semester").Order("semester DESC").Find(&courses).Pluck("Semester", &semesters).Error; err != nil {
		return nil, err
	}
	return semesters, nil
}

func (cs *CourseService) GetSectionByClassID(classID string) ([]model.Class, error) {
	var sections []model.Class
    if err := cs.db.
        Model(&model.Class{}).
        Preload("Course", func(db *gorm.DB) *gorm.DB {
            return db.Select("ID","course_code")
        }).
        Find(&sections, classID).Error; err != nil {
        return nil, err
    }
    return sections, nil
}

func (cs *CourseService) CreateCourse(course *model.Course) error {
	if err := cs.db.Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CourseService) FindCourseByID(id string) (*model.Course, error) {
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

func (cs *CourseService) DeleteCourseByID(id string) error {
	course := model.Course{}
	if result := cs.db.Delete(&course, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete the Course")
	}

	return nil
}