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

func (cs *CourseService) GetAllDistinctSemester() ([]uint, error) {
	var courses []model.Course
	var semesters []uint
	if err := cs.db.Distinct("semester").Order("semester DESC").Find(&courses).Pluck("Semester", &semesters).Error; err != nil {
		return nil, err
	}
	return semesters, nil
}

type GetSectionByClassIDField struct {
    CourseCode string `example:"CPE313"`
    Section    string `example:"A"`
}

func (cs *CourseService) GetSectionByClassID(classID string) (*GetSectionByClassIDField, error) {
	// uses join + smart select
	var sections *GetSectionByClassIDField
    if err := cs.db.Debug().
        Model(&model.Class{}).
        Select("courses.course_code AS CourseCode, classes.section AS section").
        Joins("INNER JOIN courses ON classes.course_id = courses.id").
        First(&sections, classID).Error; err != nil {
        return nil, err
    }
    return sections, nil
	
	// no join
	// var sections *model.Class
    // if err := cs.db.
    //     Model(&model.Class{}).
    //     Preload("Course", func(db *gorm.DB) *gorm.DB {
    //         return db.Select("ID","course_code")
    //     }).
    //     First(&sections, classID).Error; err != nil {
    //     return nil, err
    // }
    // return sections, nil
}


func (cs *CourseService) CreateCourse(course *model.Course) error {
	if err := cs.db.Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CourseService) GetCourseByID(id string) (*model.Course, error) {
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