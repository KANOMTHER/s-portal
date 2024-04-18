package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func PaymentRoutes(route *gin.RouterGroup, service *service.PaymentService) {
	PaymentHandler := handlers.NewPaymentHandler(service)

	payment := route.Group("/payment")
	{
		payment.GET("/get", PaymentHandler.Dummy)
	}
}
