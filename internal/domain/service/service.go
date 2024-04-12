package service

import (
	"gorm.io/gorm"
)

type Service struct {
	FacultyService *FacultyService
	CourseService *CourseService
	ProgramService *ProgramService
	TAService *TAService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		FacultyService: NewFacultyService(db),
		CourseService: NewCourseService(db),
		ProgramService: NewProgramService(db),
		TAService: NewTAService(db),
	}
}