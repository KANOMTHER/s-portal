package service

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type StudentService struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{
		db: db,
	}
}

func (ss *StudentService) CreateStudent(student *model.CreateStudentFields) (int, error) {
	Age := &model.AgingHandler{Student: student}
	Advior := &model.AdvisorHandler{Db: ss.db, Student: student}
	Pop := &model.PopulationHandler{Db: ss.db, Student: student}
	Create := &model.CreateStudentHandler{Db: ss.db, Student: student}

	Age.SetNext(Advior)
	Advior.SetNext(Pop)
	Pop.SetNext(Create)

	return Age.HandleRequest()
}

func (ss *StudentService) GetDistinctYears() ([]uint, error) {
	var years []uint
	if err := ss.db.Model(&model.Student{}).
		Select("DISTINCT CAST(FLOOR(ID / 1000000000) AS UNSIGNED) AS year").
		Order("year DESC").
		Pluck("year", &years).Error; err != nil {
		return nil, err
	}
	return years, nil
}

func (ss *StudentService) GetStudentsIDByYear(year string) ([]uint, error) {
	var students []uint
	if err := ss.db.Model(&model.Student{}).Where("CAST(ID / 1000000000 AS UNSIGNED) = ?", year).Order("ID").Pluck("ID", &students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (ss *StudentService) GetStudentByID(id string) (*model.Student, error) {
	var student *model.Student
	if err := ss.db.Preload("Program.Faculty").Preload("Advisor.Faculty").First(&student, id).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (ss *StudentService) UpdateStudentByID(context *gin.Context, id string, authSer *AuthService) (untyped int, err error) {
	user, err := authSer.GetContextUser(context)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusNotFound, nil
	}

	// Create the context
	updateContext := &model.UpdateContext{}

	// Set the strategy based on the user role
	if user.Role == "Admin" {
		updateContext.SetStrategy(&model.AdminUpdateStrategy{StudentData: model.StudentData{Db: ss.db}})
	} else if user.Role == "student" {
		updateContext.SetStrategy(&model.StudentUpdateStrategy{StudentData: model.StudentData{Db: ss.db}})
	}

	// Delegate the update operation to the selected strategy
	status := 0
	if status, err := updateContext.UpdateStudent(context, id); err != nil {
		fmt.Println("Error updating student:", err)
		return status, err
	}

	return status, nil

}

func (ss *StudentService) IsTA(id string) (*uint, error) {
	var ID *uint
	if err := ss.db.Debug().
		Model(&model.TA{}).
		Where("student_id = ?", id).
		Select("ID").
		Scan(&ID).
		Error; err != nil {
		return nil, err
	}

	return ID, nil
}
