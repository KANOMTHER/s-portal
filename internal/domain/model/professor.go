package model

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model `swaggerignore:"true"`
	ID            uint `gorm:"primaryKey" example:"1"`
	FName         string `example:"John"`
	LName         string `example:"Doe"`
	Email 		  string `example:"Juwan98@example.net"`
	Phone         string `example:"744-512-3072"`
	Position      string `example:"Professor"`
	FacultyID     uint `example:"1"`
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