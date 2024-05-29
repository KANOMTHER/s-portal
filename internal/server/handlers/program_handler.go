package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type ProgramHandler struct {
	programService *service.ProgramService
	authService    *service.AuthService
}

func NewProgramHandler(programService *service.ProgramService, authService *service.AuthService) *ProgramHandler {
	return &ProgramHandler{
		programService: programService,
		authService:    authService,
	}
}

// @Summary		GetAllPrograms
// @Description	get all programs
// @Tags			Program
// @Accept			json
// @Produce		json
// @Success		200	{array}		model.Program
// @Failure		404	{object}	string
// @Router			/program [GET]
func (h *ProgramHandler) GetAllPrograms(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

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

// @Summary		CreateProgram
// @Description	create a new program
// @Tags			Program
// @Accept			json
// @Produce		json
// @Param			program	body		model.Program	true	"Program object"
// @Success		200		{object}	string
// @Failure		400		{object}	string
// @Router			/program [POST]
func (h *ProgramHandler) CreateProgram(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

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

// @Summary		GetProgramByID
// @Description	get a program by id
// @Tags			Program
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"program id"
// @Success		200	{object}	model.Program
// @Failure		404	{object}	string
// @Router			/program/{id} [GET]
func (h *ProgramHandler) GetProgramByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	program, err := h.programService.GetProgramByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": program,
	})
}

// @Summary		GetProgramByFacultyID
// @Description	get a program by faculty_id
// @Tags			Program
// @Accept			json
// @Produce		json
// @Param			faculty_id	query		string	true	"faculty id"
// @Success		200			{array}		service.GetProgramByFacultyIDField
// @Failure		404			{object}	string
// @Router			/program/faculty [GET]
func (h *ProgramHandler) GetProgramByFacultyID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	fauclty_id := context.Query("faculty_id")
	program, err := h.programService.GetProgramByFacultyID(fauclty_id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return programs
	context.JSON(200, gin.H{
		"message": program,
	})
}

// @Summary		UpdateProgramByID
// @Description	update a program by id
// @Tags			Program
// @Accept			json
// @Produce		json
// @Param			id		path		string			true	"program id"
// @Param			Program	body		model.Program	true	"Program object"
// @Success		200		{object}	string
// @Failure		404		{object}	string
// @Router			/program/update/{id} [PUT]
func (h *ProgramHandler) UpdateProgramByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	err := h.programService.UpdateProgramByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Program updated successfully",
	})
}

// @Summary		DeleteProgramByID
// @Description	delete a program by id
// @Tags			Program
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"program id"
// @Success		200	{object}	string
// @Failure		404	{object}	string
// @Router			/program/delete/{id} [DELETE]
func (h *ProgramHandler) DeleteProgramByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	err := h.programService.DeleteProgramByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Program deleted successfully",
	})
}
