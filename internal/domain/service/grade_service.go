package service

import (
	"s-portal/internal/domain/model"

	"gorm.io/gorm"
)

type GradeService struct {
	db *gorm.DB
}

func NewGradeService(db *gorm.DB) *GradeService {
	return &GradeService{
		db: db,
	}
}

func (g *GradeService) InitialGradeMultiple(year uint, semester uint) error {
	type student_class_register struct {
		PaymentID uint
		StudentID uint
		ClassID   uint
	}

	var class_payments []student_class_register
	if err := g.db.Model(model.ClassRegister{}).Joins("left join payments on class_registers.payment_id = payments.id").Where("payments.year = ? AND payments.semester = ?", year, semester).Select("class_registers.payment_id, payments.student_id, class_registers.class_id").Find(&class_payments).Error; err != nil {
		return err
	}

	gradeHub := model.GradeHub{}
	for i := 0; i < len(class_payments); i++ {
		newGrade := gradeHub.GradeFromRegisterAdapter(class_payments[i].ClassID, class_payments[i].StudentID, 0)
		if err := g.db.Where(newGrade).FirstOrCreate(&newGrade).Error; err != nil {
			return err
		}
	}

	return nil
}

func (g *GradeService) GetStudentIDFromClassID(classID uint) ([]model.Student, error) {
	var studentID []model.Student

	if err := g.db.Model(model.Grade{}).Joins("left join students on students.id = grades.student_id").Where("grades.class_id = ?", classID).Select("students.id as id, students.f_name as f_name, students.l_name as l_name").Find(&studentID).Error; err != nil {
		return nil, err
	}

	return studentID, nil
}

func (g *GradeService) EditGradeMultiple(grade []model.Grade) error {
	for i := 0; i < len(grade); i++ {
		if err := g.db.Model(model.Grade{}).Where("student_id = ? AND class_id = ?", grade[i].StudentID, grade[i].ClassID).Update("grade", grade[i].Grade).Error; err != nil {
			return err
		}
	}
	return nil
}
