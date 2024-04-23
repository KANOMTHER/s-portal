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

type CreateStudentFields struct {
	ID        uint      `swaggerignore:"true"`
	ProgramID uint      `example:"1" binding:"required"`
	Degree    string    `example:"Bachelor"`
	Year      int       `example:"2021"`
	FName     string    `example:"Nontawat"`
	LName     string    `example:"Kunlayawuttipong"`
	DOB       time.Time `example:"2002-12-18T00:00:00Z" binding:"required"`
	Entered   time.Time `swaggerignore:"true"`
	AdvisorID uint      `example:"1" binding:"required"`
}

type UpdateStudentFields struct {
	FName     string     `example:"Nontawat"`
	LName     string     `example:"Kunlayawuttipong"`
	Graduated *time.Time `example:"2024-04-16T00:00:00Z"`
	Email     string     `example:"example@hotmail.com"`
	Phone     string     `example:"0812345678"`
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
