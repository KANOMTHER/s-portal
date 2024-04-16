package service

import (

	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type StudentService struct {
	db *gorm.DB
}

type CreateStudentFields struct {
	ID       uint      `example:"64070501093"`
	ProgramID uint      `example:"1"`
	Degree   string    `example:"Bachelor"`
	Year     int       `example:"2021"`
	FName    string    `example:"Nontawat"`
	LName    string    `example:"Kunlayawuttipong"`
	DOB      time.Time `example:"2002-12-18T00:00:00Z"`
	Entered  time.Time `example:"2024-04-16T00:00:00Z"`
	AdvisorID uint      `example:"1"`
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{
		db: db,
	}
}

func (ss *StudentService) createNewStudentId(student *CreateStudentFields) (*uint, error) {
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
	if err := ss.db.Model(&model.Program{}).Where("ID = ?", student.ProgramID).Pluck("Prefix", &program_prefix).Error; err != nil {
		return nil, err
	}
	program, err := strconv.ParseUint(program_prefix, 10, 64)
	if err != nil {
		return nil, err
	}

	// mask := 64 0705010 00
	// max_mask := 64 0705011 99
	mask :=  uint(year) + uint(program*100)
	max_mask := mask + 199

	var max_id *uint
	if err := ss.db.Model(&model.Student{}).
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
	
	// create new id for new student (max+1)
	new_id := *max_id +1
	return &new_id, nil
}

func (ss *StudentService) CreateStudent(student *CreateStudentFields) (error) {
	var err error
	newID, err := ss.createNewStudentId(student)
	if err != nil {
		return err
	}
	student.ID = *newID
	student.Entered = time.Now()

	// create new student
	if result := ss.db.Where("ID = ?", student.ID).FirstOrCreate(&model.Student{}, &student); result.Error != nil {
		if result.RowsAffected == 0 {
			return fmt.Errorf("unable to create student because this program is full %v,\n error msg: %v", student.ID, result.Error.Error())
		}
		return result.Error
	}

	return nil
}
