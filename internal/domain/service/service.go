package service

import (
	"gorm.io/gorm"
)

type Service struct {
	StudentService       *StudentService
	FacultyService       *FacultyService
	CourseService        *CourseService
	ProgramService       *ProgramService
	TimeTableService     *TimeTableService
	ClassService         *ClassService
	ProfessorService     *ProfessorService
	InstructorService    *InstructorService
	ClassRegisterService *ClassRegisterService
	PaymentService       *PaymentService
	TAService            *TAService
	AuthService          *AuthService
	GradeService         *GradeService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		StudentService:       NewStudentService(db),
		FacultyService:       NewFacultyService(db),
		CourseService:        NewCourseService(db),
		ProgramService:       NewProgramService(db),
		TimeTableService:     NewTimeTableService(db),
		ClassService:         NewClassService(db),
		ProfessorService:     NewProfessorService(db),
		InstructorService:    NewInstructorService(db),
		ClassRegisterService: NewClassRegisterService(db),
		PaymentService:       NewPaymentService(db),
		TAService:            NewTAService(db),
		AuthService:          NewAuthService(db),
		GradeService:         NewGradeService(db),
	}
}
