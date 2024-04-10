package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
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

func (h *ProgramHandler) CreateProgram(context *gin.Context) {
	program := model.Program{}

	if err := context.ShouldBindJSON(&program); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.programService.CreateProgram(&program); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Program created successfully",
	})
}

func (h *ProgramHandler) GetProgramByID(context *gin.Context) {
	id := context.Param("id")
	program, err := h.programService.GetProgramByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": program,
	})
}

func (h *ProgramHandler) UpdateProgramByID(context *gin.Context) {
	id := context.Param("id")
	err := h.programService.UpdateProgramByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Program updated successfully",
	})
}
