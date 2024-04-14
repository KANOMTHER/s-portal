package model

import (
	"time"

	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model `swaggerignore:"true"`
	ID		  uint `gorm:"primaryKey" example:"1"`
	ClassID   uint `example:"1" binding:"required"`
	Class     Class
	Day       time.Weekday `example:"0"`
	StartTime time.Time `example:"2021-08-01T08:00:00Z"`
	EndTime   time.Time `example:"2021-08-01T09:00:00Z"`
	Classroom string `example:"CPE1102"`
	ClassType string `example:"Lecture"`
}
