package model

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	ClassID      Class     `gorm:"primaryKey"`
	InstructorID Professor `gorm:"primaryKey"`
}
