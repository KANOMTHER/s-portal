package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID  uint   `gorm:"primaryKey"`
	PWD string
	Role string
}