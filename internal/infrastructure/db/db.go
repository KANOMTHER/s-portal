package db

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/joho/godotenv"
)

func connectDB() *gorm.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PWD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Println("Connected to database")
	log.Default().Println(db.Name())

	return db;
}
