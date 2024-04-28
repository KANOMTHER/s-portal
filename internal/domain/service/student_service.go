package service

import (
	"strconv"
	"time"

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

type UpdateStudentFields_ADMIN struct {
	FName     string     `example:"Nontawat"`
	LName     string     `example:"Kunlayawuttipong"`
	Graduated *time.Time `example:"2024-04-16T00:00:00Z"`
	Email     string     `example:"example@hotmail.com"`
	Phone     string     `example:"0812345678"`
}

func (ss *StudentService) UpdateStudentByID_ADMIN(context *gin.Context, id string) error {
	student := UpdateStudentFields_ADMIN{}
	if err := ss.db.Model(&model.Student{}).First(&student, id).Error; err != nil {
		return err
	}
		
	if err := context.ShouldBindJSON(&student); err != nil {
		return err
	}

	if err := ss.db.Model(&model.Student{}).
		Where("ID = ?", id).
		Updates(map[string]interface{}{
			"FName":     student.FName,
			"LName":     student.LName,
			"Graduated": student.Graduated,
			"Email":     student.Email,
			"Phone":     student.Phone,
		}).Error; err != nil {
		return err
	}
	return nil
}

type UpdateStudentFields_STUDENT struct {
	Email     string     `example:"example@hotmail.com"`
	Phone     string     `example:"0812345678"`
}

func (ss *StudentService) UpdateStudentFields_STUDENT(context *gin.Context, id string) error {
	student := UpdateStudentFields_STUDENT{}
	if err := ss.db.Model(&model.Student{}).First(&student, id).Error; err != nil {
		return err
	}
		
	if err := context.ShouldBindJSON(&student); err != nil {
		return err
	}

	if err := ss.db.Model(&model.Student{}).
		Where("ID = ?", id).
		Updates(map[string]interface{}{
			"Email":     student.Email,
			"Phone":     student.Phone,
		}).Error; err != nil {
		return err
	}
	return nil
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
