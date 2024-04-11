package model

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	CourseID   uint
	Course     Course
	Section    string
}
