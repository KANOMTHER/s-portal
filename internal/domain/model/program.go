package model

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	ProgramID    uint    `gorm:"primaryKey"`
	Faculty      Faculty
	ProgramName  string
	Detail       string
	PricePerTerm float64
	TotalTerms   int
	Prefix       string
}
