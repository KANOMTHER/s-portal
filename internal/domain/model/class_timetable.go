package model

import (
	"time"

	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model
	ID		  uint `gorm:"primaryKey"`
	ClassID   uint
	Class     Class
	Day       time.Weekday
	StartTime time.Time
	EndTime   time.Time
	Classroom string
	ClassType string
}
