package service

import (
	"fmt"
	"strconv"
	"time"

	"net/http"

	"s-portal/internal/domain/model"

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

// ------------------------------------------------------------
// AgingHandler handles student registration.
type AgingHandler struct {
	BaseHandler
	student *CreateStudentFields
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

	if ah.calculateAge(ah.student.DOB) < 10 {
		return http.StatusBadRequest, fmt.Errorf("student must be at least 10 years old")
	}

	// Call the next handler in the chain.
	if ah.nextHandler != nil {
		return ah.nextHandler.HandleRequest()
	}
	return http.StatusOK, nil
}

// ------------------------------------------------------------
// PopulationHandler handles student registration.
type PopulationHandler struct {
	BaseHandler
	db      *gorm.DB
	student *CreateStudentFields
}

func (ph *PopulationHandler) HandleRequest() (untyped int, err error) {
	ss := NewStudentService(ph.db)
	max_id, err := ss.getMaxStudentId(ph.student, ph.db)
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
// DatabaseHandler handles student registration.
type DatabaseHandler struct {
	BaseHandler
	db      *gorm.DB
	student *CreateStudentFields
}

func (rh *DatabaseHandler) createNewStudentRecord() (untyped int, err error) {
	// add in student table
	if result := rh.db.FirstOrCreate(&model.Student{}, &rh.student); result.Error != nil {
		return http.StatusBadRequest, result.Error
	}

	// add in user table
	user := model.User{ID: rh.student.ID, PWD: strconv.FormatUint(uint64(rh.student.ID), 10), Role: "student"}
	if err := rh.db.Create(&user).Error; err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (rh *DatabaseHandler) HandleRequest() (untyped int, err error) {
	ss := NewStudentService(rh.db)
	maxID, err := ss.getMaxStudentId(rh.student, rh.db)
	if err != nil {
		return http.StatusBadRequest, err
	}

	rh.student.ID = *maxID + 1
	rh.student.Entered = time.Now()
	rh.student.Year = 1 // change year to academic year

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
