package handlers

import (
	// "s-portal/internal/domain/model"

	"s-portal/internal/domain/model"
	"s-portal/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GradeHandler struct {
	gradeService *service.GradeService
	authService  *service.AuthService
}

func NewGradeHandler(gradeService *service.GradeService, authService *service.AuthService) *GradeHandler {
	return &GradeHandler{
		gradeService: gradeService,
		authService:  authService,
	}
}

func (g *GradeHandler) InitialAll(context *gin.Context) {
	type sem_year struct {
		Semester uint `json:"semester"`
		Year     uint `json:"year"`
	}

	var data sem_year

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := g.gradeService.InitialGradeMultiple(data.Year, data.Semester)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"message": "init complete",
		})
	}
}

func (g *GradeHandler) GetStudentFromClassID(context *gin.Context) {
	id := context.Param("id")
	classID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	studentID, err := g.gradeService.GetStudentIDFromClassID(uint(classID))
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"message": studentID,
		})
	}
}

func (g *GradeHandler) EditGradeMultiple(context *gin.Context) {
	id := context.Param("id")
	classID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		context.JSON(401, gin.H{
			"message": err.Error(),
		})

		return
	}

	type MultipleGrade struct {
		StudentID []uint    `json:"StudentID"`
		Grade     []float32 `json:"Grade"`
	}
	var gradeInfo MultipleGrade
	if err := context.ShouldBindJSON(&gradeInfo); err != nil {
		context.JSON(402, gin.H{
			"message": err.Error(),
		})

		return
	}

	var grades []model.Grade
	gradeHub := model.GradeHub{}
	for i := 0; i < len(gradeInfo.StudentID); i++ {
		newGrade := gradeHub.GradeFromRegisterAdapter(uint(classID), gradeInfo.StudentID[i], gradeInfo.Grade[i])
		grades = append(grades, newGrade)
	}

	err2 := g.gradeService.EditGradeMultiple(grades)
	if err2 != nil {
		context.JSON(403, gin.H{
			"message": err2.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"message": "Edit success",
		})
	}
}
