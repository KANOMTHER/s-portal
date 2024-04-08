package model

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	FacultyID    uint
	Faculty      Faculty
	ProgramName  string
	Detail       string
	PricePerTerm float64
	Prefix       string
}
