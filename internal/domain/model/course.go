package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID	uint `gorm:"primaryKey"`
	CourseCode  string
	CourseName  string
	Detail 		string
	Credit      float32
	Semester	int
	Year		int
}