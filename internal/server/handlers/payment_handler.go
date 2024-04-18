package handlers

import (
	//"fmt"

	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) Dummy(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "",
	})
}
