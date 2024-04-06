package model

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Section string
	Course  Course
}
