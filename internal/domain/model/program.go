package model

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model `swaggerignore:"true"`
	ID           uint `gorm:"primaryKey"`
	FacultyID    uint `example:"1" binding:"required"`
	Faculty      Faculty
	ProgramName  string `example:"Regular"`
	Detail       string `example:"Regular Program"`
	PricePerTerm float64 `example:"10000"`
	Prefix       string `example:"7050"`
}
