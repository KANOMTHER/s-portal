package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
