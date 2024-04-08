package main

import (
	"fmt"

	"s-portal/internal/domain/service"
	"s-portal/internal/infrastructure/db"
	"s-portal/internal/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database := db.Connect()

	// Initialize service
	services := service.NewService(database)

	// Initialize routes
	router := gin.Default()
	router = routes.InitializeRoutes(router, services)

	fmt.Println("Server started at port 3000")
	router.Run(":3000")
}