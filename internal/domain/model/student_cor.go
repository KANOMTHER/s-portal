package model

import (
	"fmt"

	"strconv"
	"time"

	"net/http"

	"gorm.io/gorm"
)

// Handler defines the interface for handling registration requests.
type Handler interface {
	HandleRequest() (untyped int, err error)
	SetNext(handler Handler)
}

// BaseHandler provides a default implementation for setting the next handler.
type BaseHandler struct {
	nextHandler Handler
}

// SetNext sets the next handler in the chain.
func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

type CreateStudentFields struct {
	ID        uint      `swaggerignore:"true"`
	ProgramID uint      `example:"1" binding:"required"`
	Degree    string    `example:"Bachelor"`
	Year      int       `example:"2021"`
	FName     string    `example:"Nontawat"`
	LName     string    `example:"Kunlayawuttipong"`
	DOB       time.Time `example:"2002-12-18T00:00:00Z" binding:"required"`
	Entered   time.Time `swaggerignore:"true"`
	AdvisorID uint      `example:"1" binding:"required"`
}

// ------------------------------------------------------------
type BaseStudentFunc struct {
}

func (bsf *BaseStudentFunc) getMaxStudentId(student *CreateStudentFields, Db *gorm.DB) (*uint, error) {
	/*
		64 0705010 93
		----------------
		64 		year
		0705010 program_prefix
		093 	max_id + 1
	*/

	// year = 64 0000000 00
	year := (student.Year - 1957) * 1000000000

	// program = 0705010
	var program_prefix string
	if err := Db.Model(&Program{}).Where("ID = ?", student.ProgramID).Pluck("Prefix", &program_prefix).Error; err != nil {
		return nil, err
	}

	if program_prefix == "" {
		return nil, fmt.Errorf("program not found")
	}

	program, err := strconv.ParseUint(program_prefix, 10, 64)
	if err != nil {
		return nil, err
	}

	// mask := 64 0705010 00
	// max_mask := 64 0705011 99
	RANGE_PROGRAM := uint(199)
	mask := uint(year) + uint(program*100)
	max_mask := mask + RANGE_PROGRAM

	var max_id *uint
	if err := Db.Model(&Student{}).
		Where("ID > ? AND ID < ?", mask, max_mask).
		Select("MAX(id)").
		Scan(&max_id).
		Error; err != nil {
		return nil, err
	}

	// if this is the first Student of the program, assign the mask
	if max_id == nil {
		max_id = &mask
	}
	return max_id, nil
}

// AgingHandler handles Student registration.

type AgingHandler struct {
	BaseHandler
	BaseStudentFunc
	Student *CreateStudentFields
}

func (ah *AgingHandler) calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Check if the birthday has occurred this year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}

func (ah *AgingHandler) HandleRequest() (untyped int, err error) {

	age := ah.calculateAge(ah.Student.DOB)
	// Validate the DOB syntax
	if ah.Student.DOB.IsZero() || ah.Student.DOB.After(time.Now()) {
		return http.StatusBadRequest, fmt.Errorf("invalid date of birth")
	}
	// Check if the Student is at least 10 years old.
	if age < 10 {
		return http.StatusBadRequest, fmt.Errorf("student must be at least 10 years old")
	}
	// Check if the Student is at most 122 years old.
	if age > 122 {
		return http.StatusBadRequest, fmt.Errorf("doesn't seem like a valid age, please check again. (current age: %d)", age)
	}

	// Call the next handler in the chain.
	if ah.nextHandler != nil {
		return ah.nextHandler.HandleRequest()
	}
	return http.StatusOK, nil
}

// ------------------------------------------------------------
// AdvisorHandler handles Student registration.

type AdvisorHandler struct {
	BaseHandler
	Db      *gorm.DB
	Student *CreateStudentFields
}

func (ah *AdvisorHandler) HandleRequest() (untyped int, err error) {
	// Check if the Advisor exists.
	var count int64
	if err := ah.Db.Model(&Professor{}).Where("ID = ?", ah.Student.AdvisorID).Count(&count).Error; err != nil {
		return http.StatusBadRequest, err
	}
	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("advisor not found")
	}

	// Call the next handler in the chain.
	if ah.nextHandler != nil {
		return ah.nextHandler.HandleRequest()
	}
	return http.StatusOK, nil
}

// ------------------------------------------------------------
// PopulationHandler handles Student registration.
type PopulationHandler struct {
	BaseHandler
	BaseStudentFunc
	Db      *gorm.DB
	Student *CreateStudentFields
}

func (ph *PopulationHandler) HandleRequest() (untyped int, err error) {
	max_id, err := ph.getMaxStudentId(ph.Student, ph.Db)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// 64 0705010 93
	// order = max_id % 100 // order = 93
	// mask = max_id - order // mask = 64 0705010 00
	mask := *max_id - (*max_id % 100)
	new_id := *max_id + 1
	MAX_STUDENT := uint(199)
	if (new_id - (mask + 1) + 1) > MAX_STUDENT {
		return http.StatusBadRequest, fmt.Errorf("program is full")
	}

	// Call the next handler in the chain.
	if ph.nextHandler != nil {
		return ph.nextHandler.HandleRequest()
	}
	return http.StatusOK, nil
}

// ------------------------------------------------------------
// CreateStudentHandler handles Student registration.
type CreateStudentHandler struct {
	BaseHandler
	BaseStudentFunc
	Db      *gorm.DB
	Student *CreateStudentFields
}

func (rh *CreateStudentHandler) createNewStudentRecord() (untyped int, err error) {
	// add in Student table
	if result := rh.Db.FirstOrCreate(&Student{}, &rh.Student); result.Error != nil {
		return http.StatusBadRequest, result.Error
	}

	// add in user table
	user := NewUser(rh.Student.ID, RoleStudent)
	if err := rh.Db.Create(&user).Error; err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (rh *CreateStudentHandler) HandleRequest() (untyped int, err error) {
	maxID, err := rh.getMaxStudentId(rh.Student, rh.Db)
	if err != nil {
		return http.StatusBadRequest, err
	}

	rh.Student.ID = *maxID + 1
	rh.Student.Entered = time.Now()
	rh.Student.Year = 1 // change year to academic year

	status, err := rh.createNewStudentRecord()
	if err != nil {
		return status, err
	}

	// Call the next handler in the chain.
	if rh.nextHandler != nil {
		return rh.nextHandler.HandleRequest()
	}

	return status, nil
}
