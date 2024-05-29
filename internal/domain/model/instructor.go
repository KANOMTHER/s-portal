package model

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model  `swaggerignore:"true"`
	ID          uint `gorm:"primaryKey" example:"1"`
	ClassID     uint `example:"1"`
	Class       Class
	ProfessorID uint `example:"1"`
	Professor   Professor
}
