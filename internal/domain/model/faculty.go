package model

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model `swaggerignore:"true"`
	ID         uint `gorm:"primaryKey" example:"1"`
	Major      string `example:"Engineering"`
	Department string `example:"Computer Engineering"`
}
