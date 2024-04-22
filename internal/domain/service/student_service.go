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

type CreateStudentFields struct {
	ID        uint      `swaggerignore:"true"`
	ProgramID uint      `example:"1"`
	Degree    string    `example:"Bachelor"`
	Year      int       `example:"2021"`
	FName     string    `example:"Nontawat"`
	LName     string    `example:"Kunlayawuttipong"`
	DOB       time.Time `example:"2002-12-18T00:00:00Z"`
	Entered   time.Time `example:"2024-04-16T00:00:00Z"`
	AdvisorID uint      `example:"1"`
}

func (ss *StudentService) getMaxStudentId(student *CreateStudentFields, db *gorm.DB) (*uint, error) {
	/*
		64 0705010 93
		----------------
		64 		year
		0705010 program_prefix
		093 	max_id + 1
	*/

	// year = 64 0000000 00
	year := (student.Year - 1957) * 1000000000

	// program = 0705010
	var program_prefix string
	if err := db.Model(&model.Program{}).Where("ID = ?", student.ProgramID).Pluck("Prefix", &program_prefix).Error; err != nil {
		return nil, err
	}
	program, err := strconv.ParseUint(program_prefix, 10, 64)
	if err != nil {
		return nil, err
	}

	// mask := 64 0705010 00
	// max_mask := 64 0705011 99
	RANGE_PROGRAM := uint(199)
	mask := uint(year) + uint(program*100)
	max_mask := mask + RANGE_PROGRAM

	var max_id *uint
	if err := db.Model(&model.Student{}).
		Where("ID > ? AND ID < ?", mask, max_mask).
		Select("MAX(id)").
		Scan(&max_id).
		Error; err != nil {
		return nil, err
	}

	// if this is the first student of the program, assign the mask
	if max_id == nil {
		max_id = &mask
	}
	return max_id, nil
}

func (ss *StudentService) CreateStudent(student *CreateStudentFields) (int, error) {
	Age := &AgingHandler{student: student}
	Pop := &PopulationHandler{db: ss.db, student: student}
	Regis := &DatabaseHandler{db: ss.db, student: student}

	Age.SetNext(Pop)
	Pop.SetNext(Regis)

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

type UpdateStudentFields struct {
	FName     string     `example:"Nontawat"`
	LName     string     `example:"Kunlayawuttipong"`
	Graduated *time.Time `example:"2024-04-16T00:00:00Z"`
	Email     string     `example:"example@hotmail.com"`
	Phone     string     `example:"0812345678"`
}

func (ss *StudentService) UpdateStudentByID(context *gin.Context, id string) error {
	student := UpdateStudentFields{}
	if err := ss.db.First(&model.Student{}, id).Error; err != nil {
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

