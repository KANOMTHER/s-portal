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
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		FacultyService: NewFacultyService(db),
		CourseService: NewCourseService(db),
		ProgramService: NewProgramService(db),
		TimeTableService: NewTimeTableService(db),
		ClassService: NewClassService(db),
	}
}