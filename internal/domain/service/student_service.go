package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type StudentService struct {
	db               *gorm.DB
	strategyRegistry *model.StrategyRegistry
}

func NewStudentService(db *gorm.DB) *StudentService {
	registry := model.NewStrategyRegistry()
	var roles []model.User
	if err := db.Distinct("role").Find(&roles).Error; err != nil {
		fmt.Println("Error getting roles while creating student service:", err)
		return nil
	}

	// Register strategies for each role
	for _, role := range roles {
		switch role.Role {
		case "Admin":
			registry.Register(role.Role, &model.AdminUpdateStrategy{StudentData: model.StudentData{Db: db}})
		case "student":
			registry.Register(role.Role, &model.StudentUpdateStrategy{StudentData: model.StudentData{Db: db}})
		// Add more cases as needed for different roles

		default:
			fmt.Println("No strategy defined for your role")
		}
	}

	return &StudentService{
		db:               db,
		strategyRegistry: registry,
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

	// Get the appropriate strategy from the registry
	strategy, err := ss.strategyRegistry.GetStrategy(user.Role)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Create the context with the strategy
	updateContext := &model.UpdateContext{}
	updateContext.SetStrategy(strategy)

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

func (ss *StudentService) GetStudentSchedule(context *gin.Context) ([]GetTimetableByClassIDField, error) {
	ps := NewClassRegisterService(ss.db)
	ts := NewTimeTableService(ss.db)

	register_classes, retErr := ps.GetRegisterClassByID(context)
	if retErr != nil {
		return nil, retErr
	}

	student_timeTable := []GetTimetableByClassIDField{}

	for i := 0; i < len(register_classes); i++ {
		class_timetable, retErr := ts.GetTimetableByClassID(strconv.FormatUint(uint64(register_classes[i].ClassID), 10))
		if retErr != nil {
			return nil, retErr
		}

		for j := 0; j < len(class_timetable); j++ {
			student_timeTable = append(student_timeTable, class_timetable[j])
		}
	}

	return student_timeTable, nil
}
