package service

import (
	"gorm.io/gorm"
)

type Service struct {
	StudentService *StudentService
	FacultyService *FacultyService
	CourseService *CourseService
	ProgramService *ProgramService
	TimeTableService *TimeTableService
	ClassService *ClassService
	ProfessorService *ProfessorService
	InstructorService *InstructorService
  TAService *TAService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		StudentService: NewStudentService(db),
		FacultyService: NewFacultyService(db),
		CourseService: NewCourseService(db),
		ProgramService: NewProgramService(db),
		TimeTableService: NewTimeTableService(db),
		ClassService: NewClassService(db),
		ProfessorService: NewProfessorService(db),
		InstructorService: NewInstructorService(db),
    TAService: NewTAService(db),
	}
}
