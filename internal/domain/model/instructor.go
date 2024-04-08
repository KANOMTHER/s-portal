package model

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	ClassID    uint
	Class      Class
	ProfessorID uint
	Professor  Professor
}
