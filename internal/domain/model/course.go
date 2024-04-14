package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model `swaggerignore:"true"`
	ID	uint `gorm:"primaryKey" example:"1"`
	CourseCode  string `example:"CPE313"`
	CourseName  string `example:"signals and linear systems"`
	Detail 		string `example:"Continuous and Discrete-time signals. Mathematical representation of signals, frequency-domain representation of signals, time- domain representation of systems, transform-domain representation of systems and system architecture. First order and higher order differential equations. Frequency response, Fourier analysis, Laplace transforms, and Z-transform."`
	Credit      float32 `example:"3"`
	Semester	int `example:"2"`
	Year		int `example:"3"`
}