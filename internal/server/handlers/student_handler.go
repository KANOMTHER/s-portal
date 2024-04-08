package handlers

// import (
// 	"github.com/gin-gonic/gin"

// 	"s-portal/internal/domain/model"
// 	"s-portal/internal/infrastructure/db"
// )

// GetStudentByID godoc
// @Summary Get student by ID
// @Description Get student by ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} Student
// @Router /students/{id} [get]
// func GetStudentByID(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var student model.Student
// 	db.First(&student, id)
// 	ctx.JSON(200, student)
// }