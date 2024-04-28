package model

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateStudentFields struct {
	FName     string     `example:"Nontawat"`
	LName     string     `example:"Kunlayawuttipong"`
	Graduated *time.Time `example:"2024-04-16T00:00:00Z"`
	Email     string     `example:"example@hotmail.com"`
	Phone     string     `example:"0812345678"`
}

// ------------------------------

// UpdateStrategy defines the update strategy interface.
type UpdateStrategy interface {
	UpdateStudent(context *gin.Context, id string) (untyped int, err error)
}

// ------------------------------

// UpdateContext defines the context for selecting the appropriate update strategy.
type UpdateContext struct {
	strategy UpdateStrategy
}

// SetStrategy sets the update strategy based on permissions or access level.
func (c *UpdateContext) SetStrategy(strategy UpdateStrategy) {
	c.strategy = strategy
}

// UpdateStudent delegates the update operation to the selected strategy.
func (c *UpdateContext) UpdateStudent(context *gin.Context, id string) (untyped int, err error) {
	if c.strategy == nil {
		return http.StatusInternalServerError, fmt.Errorf("no update strategy set")
	}
	return c.strategy.UpdateStudent(context, id)
}

// ------------------------------
// common function in each strategy
type StudentData struct {
	Db *gorm.DB
}

func (sd *StudentData) UpdateStudentData(context *gin.Context, id string) (student *UpdateStudentFields, untyped int, err error) {
	// Update student data
	if err := sd.Db.Model(&Student{}).First(&student, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, fmt.Errorf("student with ID %s not found", id)
		}
		return nil, http.StatusBadRequest, err
	}

	// Bind the JSON data to the student struct
	if err := context.ShouldBindJSON(&student); err != nil {
		return nil, http.StatusBadRequest, err
	}

	return student, http.StatusOK, nil
}

// ------------------------------

// AdminUpdateStrategy defines the update strategy for admin users.
type AdminUpdateStrategy struct {
	StudentData
}

// UpdateStudent updates the student record with the provided fields.
func (as *AdminUpdateStrategy) UpdateStudent(context *gin.Context, id string) (untyped int, err error) {

	// Update student data
	student, status, err := as.UpdateStudentData(context, id)
	if err != nil {
		return status, err
	}

	// Update the student record
	if err := as.Db.Model(&Student{}).
		Where("ID = ?", id).
		Updates(map[string]interface{}{
			"FName":     student.FName,
			"LName":     student.LName,
			"Graduated": student.Graduated,
			"Email":     student.Email,
			"Phone":     student.Phone,
		}).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// ------------------------------

// StudentUpdateStrategy defines the update strategy for student users.
type StudentUpdateStrategy struct {
	StudentData
}

// UpdateStudent updates the student record with the provided fields.
func (ss *StudentUpdateStrategy) UpdateStudent(context *gin.Context, id string) (untyped int, err error) {
	
	// Update student data
	student, status, err := ss.UpdateStudentData(context, id)
	if err != nil {
		return status, err
	}

	// Update the student record
	if err := ss.Db.Model(&Student{}).
		Where("ID = ?", id).
		Updates(map[string]interface{}{
			"Email": student.Email,
			"Phone": student.Phone,
		}).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
