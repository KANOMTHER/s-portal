package model

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model `swaggerignore:"true"`
	ID         uint   `gorm:"primaryKey" example:"1"`
	CourseID   uint   `example:"1" binding:"required"`
	Course     Course 
	Section    string `example:"A"`
}
