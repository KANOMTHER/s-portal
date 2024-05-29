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

func (h *TAHandler) GetClassTA(context *gin.Context) {
	ta, err := h.TAService.GetTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": ta,
	})
}

func (h *TAHandler) GetClassTAByClassID(context *gin.Context) {
	ta, err := h.TAService.GetTAByClassID(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": ta,
	})
}

func (h *TAHandler) CreateClassTA(context *gin.Context) {
	err := h.TAService.CreateTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": "create ta success",
	})
}

func (h *TAHandler) UpdateClassTA(context *gin.Context) {
	err := h.TAService.UpdateTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": "Update ta success",
	})
}

func (h *TAHandler) DeleteClassTA(context *gin.Context) {
	err := h.TAService.DeleteTA(context)
	
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message err": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": "delete ta complete",
	})
}