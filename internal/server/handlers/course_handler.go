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

func (h *CourseHandler) GetAllCourses(context *gin.Context) {
	courses, err := h.courseService.GetAllCourses()
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No courses found",
		})
	}
	// Return courses
	context.JSON(200, gin.H{
		"message": courses,
	})
}

func (h *CourseHandler) CreateCourse(context *gin.Context) {
	course := model.Course{}

	if err := context.ShouldBindJSON(&course); err != nil {
		// Handle error
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.courseService.CreateCourse(context, &course); err != nil {
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

func (h *CourseHandler) FindCourseByID(context *gin.Context) {
	id := context.Param("id")
	course, err := h.courseService.FindCourseByID(context, id)
	if err != nil {
		// Handle error
		context.JSON(404, gin.H{
			"message": "No course found",
		})
		return;
	}
	// Return courses
	context.JSON(200, gin.H{
		"message": course,
	})
}
