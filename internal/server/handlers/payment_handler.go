package handlers

import (
	//"fmt"

	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
	authService    *service.AuthService
}

func NewPaymentHandler(paymentService *service.PaymentService, authService *service.AuthService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		authService:    authService,
	}
}

func (h *PaymentHandler) Dummy(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "",
	})
}
