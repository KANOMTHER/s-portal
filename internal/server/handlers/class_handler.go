package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{
		classService: classService,
	}
}

//	@Summary		CreateClass
//	@Description	create a new class
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			class	body		model.Class	true	"Class object"
//	@Success		200		{object}	string		"Class created successfully"
//	@Failure		400		{object}	string		"some error message here (from err.Error())"
//	@Router			/class [POST]
func (h *ClassHandler) CreateClass(context *gin.Context) {
	class := model.Class{}

	if err := context.ShouldBindJSON(&class); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.classService.CreateClass(&class); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Class created successfully",
	})
}

//	@Summary		GetClassByID
//	@Description	get class search by id
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"class id"
//	@Success		200	{object}	model.Class
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/class/{id} [GET]
func (h *ClassHandler) GetClassByID(context *gin.Context) {
	id := context.Param("id")
	class, err := h.classService.GetClassByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return classs
	context.JSON(200, gin.H{
		"message": class,
	})
}

//	@Summary		GetClassByCourseID
//	@Description	get class search by course_id
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			course_id	query		string	true	"course id"
//	@Success		200			{array}		model.Class
//	@Failure		404			{object}	string	"some error message here (from err.Error())"
//	@Router			/class/course [GET]
func (h *ClassHandler) GetClassByCourseID(context *gin.Context) {
	course_id := context.Query("course_id")
	class, err := h.classService.GetClassByCourseID(course_id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return classs
	context.JSON(200, gin.H{
		"message": class,
	})
}

//	@Summary		GetClassBySemester
//	@Description	get class search by semester
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			semester	query		string	true	"1, 2, ..."
//	@Success		200			{array}		uint	"array of class_id"
//	@Failure		404			{object}	string	"some error message here (from err.Error())"
//	@Router			/class/semester [GET]
func (h *ClassHandler) GetClassBySemester(context *gin.Context) {
	semester := context.Query("semester")
	class, err := h.classService.GetClassBySemester(semester)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return classs
	context.JSON(200, gin.H{
		"message": class,
	})
}

//	@Summary		DeleteClassByID
//	@Description	delete class by class_id
//	@Tags			Class
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"class_id"
//	@Success		200	{object}	string	"Class deleted successfully"
//	@Failure		404	{object}	string	"were not able to delete the class"
//	@Router			/class/delete/{id} [DELETE]
func (h *ClassHandler) DeleteClassByID(context *gin.Context) {
	id := context.Param("id")
	err := h.classService.DeleteClassByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Class deleted successfully",
	})
}