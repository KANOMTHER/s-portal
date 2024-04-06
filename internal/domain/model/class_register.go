package model

import (
	"gorm.io/gorm"
)

type ClassRegister struct {
	gorm.Model
	Class     Class `gorm:"primaryKey"`
	StudentID Student  `gorm:"primaryKey"`
}