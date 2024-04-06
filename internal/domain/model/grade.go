package model

import (
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Class	 Class `gorm:"primaryKey"`
	Student  Student  `gorm:"primaryKey"`
	Grade    float32
}