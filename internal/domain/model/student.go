package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	ProgramID uint
	Program   Program
	Degree    string
	Year      int
	FName     string
	LName     string
	DOB       time.Time
	Entered   time.Time
	Graduated *time.Time
	Email     string
	Phone     string
	AdvisorID uint
	Advisor   Professor
}
