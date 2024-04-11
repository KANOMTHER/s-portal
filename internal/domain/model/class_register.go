package model

import (
	"gorm.io/gorm"
)

type ClassRegister struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	PaymentID uint
	Payment   Payment
	ClassID   uint
	Class     Class
}
