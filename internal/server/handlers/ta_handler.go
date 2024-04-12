package handlers

import (
	"github.com/gin-gonic/gin"

	// "s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type TAHandler struct {
	TAService *service.TAService
}

func NewTAHandler(TAService *service.TAService) *TAHandler {
	return &TAHandler{
		TAService: TAService,
	}
}

func (h *TAHandler) GetHello(context *gin.Context) {
	ta := h.TAService.GetTA()

	context.JSON(200, gin.H{
		"message": ta,
	})
}

func (h *TAHandler) CreateClassTA(context *gin.Context) {
	err := h.TAService.CreateTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": "sth",
	})
}

func (h *TAHandler) UpdateClassTA(context *gin.Context) {
	err := h.TAService.UpdateTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": "sth",
	})
}