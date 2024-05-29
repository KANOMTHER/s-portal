package model

import (
	"gorm.io/gorm"
)

type TA struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	ClassID   uint
	Class     Class
	StudentID uint
	Student   Student
}

func NewTA(classID uint, studentId uint) TA {
	return TA{
		ClassID: classID,
		//Class: class,
		StudentID: studentId,
		//Student: student,
	}
}
