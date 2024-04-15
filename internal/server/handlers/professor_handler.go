package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type ProfessorHandler struct {
	professorService *service.ProfessorService
}

func NewProfessorHandler(professorService *service.ProfessorService) *ProfessorHandler {
	return &ProfessorHandler{
		professorService: professorService,
	}
}

func (h *ProfessorHandler) GetAllProfessors(context *gin.Context) {
	professors, err := h.professorService.GetAllProfessors()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No professors found",
		})
	}
	// Return professors
	context.JSON(200, gin.H{
		"message": professors,
	})
}

func (h *ProfessorHandler) CreateProfessor(context *gin.Context) {
	professor := model.Professor{}

	if err := context.ShouldBindJSON(&professor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.professorService.CreateProfessor(&professor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Professor created successfully",
	})
}

func (h *ProfessorHandler) GetProfessorByID(context *gin.Context) {
	id := context.Param("id")
	professor, err := h.professorService.GetProfessorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return professor
	context.JSON(200, gin.H{
		"message": professor,
	})
}

func (h *ProfessorHandler) GetProfessorScheduleByID(context *gin.Context) {
	id := context.Param("id")
	schedules, err := h.professorService.GetProfessorScheduleByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return professor's schedules
	context.JSON(200, gin.H{
		"message": schedules,
	})
}

func (h *ProfessorHandler) UpdateProfessorByID(context *gin.Context) {
	id := context.Param("id")
	err := h.professorService.UpdateProfessorByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Professor updated successfully",
	})
}

func (h *ProfessorHandler) DeleteProfessorByID(context *gin.Context) {
	id := context.Param("id")
	err := h.professorService.DeleteProfessorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Professor deleted successfully",
	})
}
