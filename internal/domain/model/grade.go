package model

import (
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	ClassID   uint
	Class     Class
	StudentID uint
	Student   Student
	Grade     float32
}
