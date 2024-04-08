package routes

import (
	"github.com/gin-gonic/gin"
	// "s-portal/internal/server/handlers"
)

func StudentRoute(route *gin.RouterGroup) {
	students := route.Group("/students")
	{
		students.GET("/hi", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Yo students!",
			})
		
		})
		// route.GET("/:id", handlers.GetStudent)
		// route.POST("/", handlers.CreateStudent)
		// route.PUT("/:id", handlers.UpdateStudent)
		// route.DELETE("/:id", handlers.DeleteStudent)
	}
}