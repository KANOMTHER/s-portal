package routes

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
	"s-portal/internal/server/handlers"
)

func PaymentRoutes(route *gin.RouterGroup, service *service.PaymentService, authService *service.AuthService) {
	PaymentHandler := handlers.NewPaymentHandler(service, authService)

	payment := route.Group("/payment")
	{
		payment.GET("/get", PaymentHandler.Dummy)
	}
}
