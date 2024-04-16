package main

import (
	"fmt"

	"s-portal/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"s-portal/internal/domain/service"
	"s-portal/internal/infrastructure/db"
	"s-portal/internal/server/routes"

	"github.com/gin-gonic/gin"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	docs.SwaggerInfo.Title = "s-api"
	docs.SwaggerInfo.Description = "This is a api docs for s-api."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Connect to database
	database := db.Connect()

	// Initialize service
	services := service.NewService(database)

	// Initialize routes
	router := gin.Default()
	router = routes.InitializeRoutes(router, services)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Server started at port 3000")
	router.Run(":3000")
}
