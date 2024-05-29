package handlers

import (
	"github.com/gin-gonic/gin"

	// "s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type InitHandler struct {
	InitService *service.InitService
}

func NewInitHandler(InitService *service.InitService) *InitHandler {
	return &InitHandler{
		InitService: InitService,
	}
}

func (h *InitHandler) GoInit(context *gin.Context) {
	h.InitService.InitDatabase()

	context.JSON(200, gin.H{
		"message": "doing",
	})
}