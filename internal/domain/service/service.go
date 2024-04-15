package service

import (
	"gorm.io/gorm"
)

type Service struct {
	FacultyService *FacultyService
	CourseService *CourseService
	ProgramService *ProgramService
	ProfessorService *ProfessorService
	InstructorService *InstructorService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		FacultyService: NewFacultyService(db),
		CourseService: NewCourseService(db),
		ProgramService: NewProgramService(db),
		ProfessorService: NewProfessorService(db),
		InstructorService: NewInstructorService(db),
	}
}
