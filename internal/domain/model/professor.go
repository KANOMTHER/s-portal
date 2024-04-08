package model

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	FName         string
	LName         string
	Email 		  string
	Phone         string
	Position      string
	FacultyID     uint
	Faculty       Faculty
}

func CreateProfessor(id uint, fName string, lName string, email string, phone string, position string, facultyID uint, faculty Faculty) Professor{
	return Professor{
		ID: id,
		FName: fName,
		LName: lName,
		Email: email,
		Phone: phone,
		Position: position,
		FacultyID: facultyID,
		Faculty: faculty,
	}
}