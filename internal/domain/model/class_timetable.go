package model

import (
	"time"

	"gorm.io/gorm"
)

type ClassTimetable struct {
	gorm.Model
	Class     Class `gorm:"primaryKey"`
	No        int   `gorm:"primaryKey"`
	Day       time.Weekday
	StartTime time.Time
	EndTime   time.Time
	Classroom string
	ClassType string
}
