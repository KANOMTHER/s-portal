package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type InstructorHandler struct {
	instructorService *service.InstructorService
	authService       *service.AuthService
}

func NewInstructorHandler(instructorService *service.InstructorService, authService *service.AuthService) *InstructorHandler {
	return &InstructorHandler{
		instructorService: instructorService,
		authService:       authService,
	}
}

// @Summary		GetAllInstructors
// @Description	get all instructors
// @Tags			Instructor
// @Accept			json
// @Produce		json
// @Success		200	{array}		model.Instructor
// @Failure		404	{object}	string	"No instructors found"
// @Router			/instructor [GET]
func (h *InstructorHandler) GetAllInstructors(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	instructors, err := h.instructorService.GetAllInstructors()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No instructors found",
		})
	}
	// Return instructors
	context.JSON(200, gin.H{
		"message": instructors,
	})
}

// @Summary		CreateInstructor
// @Description	create a new instructor
// @Tags			Instructor
// @Accept			json
// @Produce		json
// @Param			instructor	body		model.Instructor	true	"Instructor object"
// @Success		200			{object}	string				"Instructor created successfully"
// @Failure		400			{object}	string				"some error message here (from err.Error())"
// @Router			/instructor [POST]
func (h *InstructorHandler) CreateInstructor(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	instructor := model.Instructor{}

	if err := context.ShouldBindJSON(&instructor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.instructorService.CreateInstructor(&instructor); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor created successfully",
	})
}

// @Summary		GetInstructorByID
// @Description	get a instructor by id
// @Tags			Instructor
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"instructor id"
// @Success		200	{object}	model.Instructor	"Instructor object"
// @Failure		404	{object}	string				"some error message here (from err.Error())"
// @Router			/instructor/{id} [GET]
func (h *InstructorHandler) GetInstructorByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	instructor, err := h.instructorService.GetInstructorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return instructors
	context.JSON(200, gin.H{
		"message": instructor,
	})
}

// @Summary		UpdateInstructorByID
// @Description	update a instructor by id
// @Tags			Instructor
// @Accept			json
// @Produce		json
// @Param			id			path		string				true	"Instructor id"
// @Param			Instructor	body		model.Instructor	true	"Instructor object"
// @Success		200			{object}	string				"Instructor updated successfully"
// @Failure		404			{object}	string				"some error message here (from err.Error())"
// @Router			/instructor/update/{id} [PUT]
func (h *InstructorHandler) UpdateInstructorByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	err := h.instructorService.UpdateInstructorByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor updated successfully",
	})
}

// @Summary		DeleteInstructorByID
// @Description	delete a instructor by id
// @Tags			Instructor
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"instructor id"
// @Success		200	{object}	string	"Instructor deleted successfully"
// @Failure		404	{object}	string	"some error message here (from err.Error())"
// @Router			/instructor/delete/{id} [DELETE]
func (h *InstructorHandler) DeleteInstructorByID(context *gin.Context) {
	if !h.authService.AssertPermission(context) {
		context.Status(401)
		return
	}

	id := context.Param("id")
	err := h.instructorService.DeleteInstructorByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Instructor deleted successfully",
	})
}
