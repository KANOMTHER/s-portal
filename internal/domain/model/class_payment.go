package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	StudentID   uint
	Student     Student
	TotalCredit float32
	CreatedAt   time.Time
	Semester    int
	Year        int
	Total       float32
	Status      string
}
