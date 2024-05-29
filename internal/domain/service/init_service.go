package service

import (
	"s-portal/internal/domain/model"
	"time"

	"gorm.io/gorm"
)

type InitService struct {
	db *gorm.DB
}

func NewInitService(db *gorm.DB) *InitService {
	return &InitService{
		db: db,
	}
}

func (is *InitService) InitDatabase() {
	faculty := model.Faculty{
		Major: "a",
		Department: "b",
	}

	program := model.Program{
		FacultyID: 1,
	}

	professor := model.Professor{
		FacultyID: 1,
	}

	student := model.Student{
		ProgramID: 1,
		DOB: time.Now(),
		Entered: time.Now(),
		AdvisorID: 1,
	}

	course := model.Course{

	}

	class := model.Class{
		CourseID: 1,
	}

	is.db.Create(&faculty)
	is.db.Create(&program)
	is.db.Create(&professor)
	is.db.Create(&student)
	is.db.Create(&course)
	is.db.Create(&class)
}