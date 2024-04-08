package model

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Major      string
	Department string
}
