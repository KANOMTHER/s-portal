package handlers

import (
	"github.com/gin-gonic/gin"


	"s-portal/internal/domain/service"
)

type ProgramHandler struct {
	programService *service.ProgramService
}

func NewProgramHandler(programService *service.ProgramService) *ProgramHandler {
	return &ProgramHandler{
		programService: programService,
	}
}

func (h *ProgramHandler) GetAllPrograms(context *gin.Context) {
	programs, err := h.programService.GetAllPrograms()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No programs found",
		})
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": programs,
	})
}
