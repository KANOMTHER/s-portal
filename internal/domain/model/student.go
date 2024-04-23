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

func CreateStudent( programID uint, program Program, degree string, year int, fName string, lName string, dob time.Time, entered time.Time, graduated *time.Time, email string, phone string, advisorID uint, advisor Professor) Student {
	return Student{
		ProgramID: programID,
		Program: program,
		Degree: degree,
		Year: year,
		FName: fName,
		LName: lName,
		DOB: dob,
		Entered: entered,
		Graduated: graduated,
		Email: email,
		Phone: phone,
		AdvisorID: advisorID,
		Advisor: advisor,
	}
}
