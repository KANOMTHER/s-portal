package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Code        string `gorm:"primaryKey"`
	Name        string
	Detail 		string
	Credit      int
	ReqCourse   []Course `gorm:"many2many:course_prerequisites;"`
	Semester	string
}