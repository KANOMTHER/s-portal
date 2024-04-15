package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type FacultyHandler struct {
	facultyService *service.FacultyService
}

func NewFacultyHandler(facultyService *service.FacultyService) *FacultyHandler {
	return &FacultyHandler{
		facultyService: facultyService,
	}
}
//	@Summary		GetAllFaculties
//	@Description	get all faculties
//	@Tags			Faculty
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Faculty	"list of faculties"
//	@Failure		404	{object}	string			"No faculties found"
//	@Router			/faculty [GET]
func (h *FacultyHandler) GetAllFaculties(c *gin.Context) {
	faculties, err := h.facultyService.GetAllFaculties()
	if err != nil {
		// Handle error
		c.JSON(404, gin.H{
			"message": "No faculties found",
		})
	}
	// Return faculties
	c.JSON(200, gin.H{
		"message": faculties,
	})
}

//	@Summary		CreateFaculty
//	@Description	create a new faculty
//	@Tags			Faculty
//	@Accept			json
//	@Produce		json
//	@Param			program	body		model.Faculty	true	"Faculty object"
//	@Success		200		{object}	string			"Faculty created successfully"
//	@Failure		400		{object}	string			"some error message here (from err.Error())"
//	@Router			/faculty [POST]
func (h *FacultyHandler) CreateFaculty(context *gin.Context) {
	faculty := model.Faculty{}

	if err := context.ShouldBindJSON(&faculty); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.facultyService.CreateFaculty(context, &faculty); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Faculty created successfully",
	})
}

//	@Summary		GetFacultyByID
//	@Description	get a faculty by id
//	@Tags			Faculty
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"faculty id"
//	@Success		200	{object}	model.Faculty
//	@Failure		404	{object}	string	"Faculty not found"
//	@Router			/faculty/{id} [GET]
func (h *FacultyHandler) GetFacultyByID(context *gin.Context) {
	id := context.Param("id")
	faculty, err := h.facultyService.GetFacultyByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "Faculty not found",
		})
		return
	}
	// Return faculty
	context.JSON(200, gin.H{
		"message": faculty,
	})
}

//	@Summary		UpdateFacultyByID
//	@Description	update a facluty by id
//	@Tags			Faculty
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Faculty id"
//	@Param			Faculty	body		model.Faculty	true	"Faculty object"
//	@Success		200		{object}	string			"Faculty updated successfully"
//	@Failure		400		{object}	string			"some error message here (from err.Error())"
//	@Router			/faculty/update/{id} [PUT]
func (h *FacultyHandler) UpdateFacultyByID(context *gin.Context) {
	id := context.Param("id")
	err := h.facultyService.UpdateFacultyByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Faculty updated successfully",
	})
}

//	@Summary		DeleteFacultyByID
//	@Description	delete a faculty by id
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"faculty id"
//	@Success		200	{object}	string	"Faculty deleted successfully"
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/faculty/delete/{id} [DELETE]
func (h *FacultyHandler) DeleteFacultyByID(context *gin.Context) {
	id := context.Param("id")
	err := h.facultyService.DeleteFacultyByID(id)
	if err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Faculty deleted successfully",
	})
}
