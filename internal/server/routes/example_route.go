package routes

import (
	"github.com/gin-gonic/gin"
)

func ExampleRoutes(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to s-portal backend",
		})
	})
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}