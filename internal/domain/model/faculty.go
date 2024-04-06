package model

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	FacultyID uint `gorm:"primaryKey"`
	MajorName string
	Department string
}