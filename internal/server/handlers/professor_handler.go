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

//	@Summary		GetAllProfessors
//	@Description	get all professors
//	@Tags			Professor
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		service.ProfessorProfile
//	@Failure		404	{object}	string	"No professors found"
//	@Router			/professor [GET]
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

//	@Summary		CreateProfessor
//	@Description	create a new professor
//	@Tags			Professor
//	@Accept			json
//	@Produce		json
//	@Param			professor	body		model.Professor	true	"Professor object"
//	@Success		200			{object}	string			"Professor created successfully"
//	@Failure		400			{object}	string			"some error message here (from err.Error())"
//	@Router			/professor [POST]
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

//	@Summary		GetProfessorByID
//	@Description	get a program by id
//	@Tags			Professor
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"professor id"
//	@Success		200	{object}	service.ProfessorProfile	"ProfessorProfile object"
//	@Failure		404	{object}	string						"some error message here (from err.Error())"
//	@Router			/professor/{id} [GET]
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

//	@Summary		UpdateProfessorByID
//	@Description	update a professor by id
//	@Tags			Professor
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string			true	"professor id"
//	@Param			Professor	body		model.Professor	true	"Program object"
//	@Success		200			{object}	string			"Professor updated successfully"
//	@Failure		404			{object}	string			"some error message here (from err.Error())"
//	@Router			/professor/update/{id} [PUT]
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

//	@Summary		DeleteProfessorByID
//	@Description	delete a professor by id
//	@Tags			Professor
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"professor id"
//	@Success		200	{object}	string	"Professor deleted successfully"
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/professor/delete/{id} [DELETE]
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
