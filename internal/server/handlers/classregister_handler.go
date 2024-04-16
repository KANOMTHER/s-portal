package handlers

import (
	//"fmt"

	"github.com/gin-gonic/gin"

	"s-portal/internal/domain/service"
)

type ClassRegisterHandler struct {
	classRegisterService *service.ClassRegisterService
}

func NewClassRegisterHandler(classRegisterService *service.ClassRegisterService) *ClassRegisterHandler {
	return &ClassRegisterHandler{
		classRegisterService: classRegisterService,
	}
}

// selection: query class -> select class -> query section -> select section -> add button (get id of class table) -> add id and student in class_register
//need
/*
 * course (query course_code(for showing in screen), course_id(store for sending next request) with year and semester in course) COMPLETE
 * class (query section(for showing in screen), class_id(store for sending next request) with course_id in class) COMPLETE
 * classregister (create record in class_register with student_id, year, semester, class_id) COMPLETE
 *		// need
 * 		* classpayment (create record in class_payment(for first added) with student_id, year, semester) CAUTION have coupling (implement at classregister)
 * 		* classpayment (query payment_id in class_payment with student_id, year, semester) CAUTION have coupling (implement at classregister)
 */

func (h *ClassRegisterHandler) GetRegisterClassByID(context *gin.Context) {
	class_register, err := h.classRegisterService.GetRegisterClassByID(context)

	if err!=nil {
		context.JSON(400, gin.H{
			"message err": err.Error(),
		})
	} else {
		// Return success message
		context.JSON(200, gin.H{
			"message": class_register,
		})
	}
}

func (h *ClassRegisterHandler) CreateRegisterClass(context *gin.Context) {
	err := h.classRegisterService.CreateRegisterClass(context)

	if err!=nil {
		context.JSON(400, gin.H{
			"message err": err.Error(),
		})
	} else {
		// Return success message
		context.JSON(200, gin.H{
			"message": "register successfully",
		})
	}
}

func (h *ClassRegisterHandler) UpdateRegisterClass(context *gin.Context) {
	err := h.classRegisterService.UpdateRegisterClass(context)

	if err!=nil {
		context.JSON(400, gin.H{
			"message err": err.Error(),
		})
	} else {
		// Return success message
		context.JSON(200, gin.H{
			"message": "update successfully",
		})
	}
}

func (h *ClassRegisterHandler) DeleteRegisterClass(context *gin.Context) {
	err := h.classRegisterService.DeleteRegisterClass(context)

	if err!=nil {
		context.JSON(400, gin.H{
			"message err": err.Error(),
		})
	} else {
		// Return success message
		context.JSON(200, gin.H{
			"message": "update successfully",
		})
	}
}