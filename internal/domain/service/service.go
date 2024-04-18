package service

import (
	"gorm.io/gorm"
)

type Service struct {
	FacultyService *FacultyService
	CourseService *CourseService
	ProgramService *ProgramService
	TimeTableService *TimeTableService
	ClassService *ClassService
	ProfessorService *ProfessorService
	InstructorService *InstructorService
	ClassRegisterService *ClassRegisterService
	PaymentService *PaymentService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		FacultyService: NewFacultyService(db),
		CourseService: NewCourseService(db),
		ProgramService: NewProgramService(db),
		TimeTableService: NewTimeTableService(db),
		ClassService: NewClassService(db),
		ProfessorService: NewProfessorService(db),
		InstructorService: NewInstructorService(db),
		ClassRegisterService: NewClassRegisterService(db),
		PaymentService: NewPaymentService(db),
	}
}
