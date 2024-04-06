package model

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	FName         string
	LName         string
	PersonalEmail string
	SchoolEmail   string
	Phone         string
	Position      string
	Faculty       Faculty
}
