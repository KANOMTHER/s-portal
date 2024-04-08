package model

import (

	"gorm.io/gorm"
)

type TA struct {
	gorm.Model
	ID		  uint `gorm:"primaryKey"`
	ClassID   uint
	Class     Class
	StudentID uint
	Student   Student
}
