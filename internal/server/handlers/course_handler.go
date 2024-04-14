package handlers

import (
	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"

)

type CourseHandler struct {
	courseService *service.CourseService
}

func NewCourseHandler(courseService *service.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

//	@Summary		GetAllCourses
//	@Description	get all courses
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Course
//	@Failure		404	{object}	string
//	@Router			/course [GET]
func (h *CourseHandler) GetAllCourses(context *gin.Context) {
	courses, err := h.courseService.GetAllCourses()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return courses
	context.JSON(200, gin.H{
		"message": courses,
	})
}

//	@Summary		GetAllDistinctSemester
//	@Description	get semester from all course [no duplicate]
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		uint
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/course/semester [GET]
func (h *CourseHandler) GetAllDistinctSemester(context *gin.Context) {
	semesters, err := h.courseService.GetAllDistinctSemester()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return semesters
	context.JSON(200, gin.H{
		"message": semesters,
	})
}

//	@Summary		GetSectionByClassID
//	@Description	get section search by class_id
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			class_id	query		string								true	"class id"
//	@Success		200			{array}		service.GetSectionByClassIDField	"CourseCode, Section"
//	@Failure		404			{object}	string								"some error message here (from err.Error())"
//	@Router			/course/section [GET]
func (h *CourseHandler) GetSectionByClassID(context *gin.Context) {
	classID := context.Query("class_id")
	sections, err := h.courseService.GetSectionByClassID(classID)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return sections
	context.JSON(200, gin.H{
		"message": sections,
	})
}

//	@Summary		CreateCourse
//	@Description	create a new course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			program	body		model.Course	false	"Course object"
//	@Success		200		{object}	string			"Course created successfully"
//	@Failure		400		{object}	string			"some error message here (from err.Error())"
//	@Router			/course [POST]
func (h *CourseHandler) CreateCourse(context *gin.Context) {
	course := model.Course{}

	if err := context.ShouldBindJSON(&course); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.courseService.CreateCourse(&course); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return success message
	context.JSON(200, gin.H{
		"message": "Course created successfully",
	})
}

//	@Summary		GetCourseByID
//	@Description	get a course by id
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"program id"
//	@Success		200	{object}	model.Course
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/course/{id} [GET]
func (h *CourseHandler) GetCourseByID(context *gin.Context) {
	id := context.Param("id")
	course, err := h.courseService.GetCourseByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return courses
	context.JSON(200, gin.H{
		"message": course,
	})
}

//	@Summary		UpdateCourseByID
//	@Description	update a course by id
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"course id"
//	@Param			Course	body		model.Course	false	"Course object"
//	@Success		200		{object}	string			"Course updated successfully"
//	@Failure		404		{object}	string			"some error message here (from err.Error())"
//	@Router			/course/update/{id} [PUT]
func (h *CourseHandler) UpdateCourseByID(context *gin.Context) {
	id := context.Param("id")
	err := h.courseService.UpdateCourseByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Course updated successfully",
	})
}

//	@Summary		DeleteCourseByID
//	@Description	delete a course by id
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"course id"
//	@Success		200	{object}	string	"Course deleted successfully"
//	@Failure		404	{object}	string	"some error message here (from err.Error())"
//	@Router			/course/delete/{id} [DELETE]
func (h *CourseHandler) DeleteCourseByID(context *gin.Context) {
	id := context.Param("id")
	err := h.courseService.DeleteCourseByID(id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": err.Error(),
		})
		return;
	}
	// Return success message
	context.JSON(200, gin.H{
		"message": "Course deleted successfully",
	})
}