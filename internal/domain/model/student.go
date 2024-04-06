package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	ProgramID     uint
	Degree        string
	Year          int
	IDCard        string
	Sex           string
	FName         string
	LName         string
	DOB           time.Time
	Entered       time.Time
	Graduated     *time.Time
	PersonalEmail string
	SchoolEmail   string
	Phone         string
	Address       string
	Advisor       Professor
}
