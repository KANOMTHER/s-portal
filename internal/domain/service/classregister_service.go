package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type ClassRegisterService struct {
	db *gorm.DB
}

func NewClassRegisterService(db *gorm.DB) *ClassRegisterService {
	return &ClassRegisterService{
		db: db,
	}
}

func (cr *ClassRegisterService) GetRegisterClassByID(context *gin.Context) ([]model.ClassRegister, error) {
	payment := model.Payment{}
	ps := NewPaymentService(cr.db)

	if err := context.ShouldBindJSON(&payment); err!=nil {
		return nil, err
	}
	
	payment, payment_err := ps.GetAndCreatePaymentByID(payment)
	if payment_err!=nil {
		return nil, payment_err
	}

	var result []model.ClassRegister

	if err := cr.db.Where("payment_id = ?", payment.ID).Preload("Class").Find(&result).Error; err!=nil {
		return nil, err
	}

	return result, nil
}

func (cr *ClassRegisterService) CreateRegisterClass(context *gin.Context) error {
	var data struct {
		StudentID uint `json:"StudentID"`
		Year      uint `json:"year"`
		Semester  uint `json:"semester"`
		ClassID   uint `json:"ClassID"`
	}
	ps := NewPaymentService(cr.db)

	if err := context.ShouldBindJSON(&data); err!=nil {
		return err
	}

	if data.ClassID == 0 || data.Semester == 0 || data.StudentID == 0 || data.Year == 0 {
		return fmt.Errorf("invalid input")
	}

	payment := model.Payment{
		StudentID: data.StudentID,
		Semester: int(data.Semester),
		Year: int(data.Year),
	}
	payment, payment_err := ps.GetAndCreatePaymentByID(payment)
	if payment_err!=nil {
		return payment_err
	}

	//require step for initial class_grade
	// -> insert after end phase

	class_register := model.ClassRegister{
		PaymentID: payment.ID,
		ClassID: data.ClassID,
	}
	//please check before insert
	if err := cr.db.Where(class_register).FirstOrCreate(&class_register).Error ; err!=nil {
		return err
	}
	ps.UpdateTotalCreditByID(class_register.PaymentID)

	return nil
}

func (cr *ClassRegisterService) UpdateRegisterClass(context *gin.Context) error { 
	var data struct{
		ID uint `json:"id"`
		ClassID uint `json:"ClassID"`
	}
	ps := NewPaymentService(cr.db)

	if err := context.ShouldBindJSON(&data); err!=nil {
		return err
	}

	class_register := model.ClassRegister{}

	if err := cr.db.First(&class_register, data.ID).Error; err!=nil {
		return err
	}

	if err := context.ShouldBindJSON(&class_register); err!=nil {
		return err
	}

	isExists := model.ClassRegister{
		PaymentID: class_register.PaymentID,
		ClassID: class_register.ClassID,
	}

	err := cr.db.Where(isExists).First(&isExists).Error
	if err!=gorm.ErrRecordNotFound {
		return fmt.Errorf("your class_id is already registed")
	}

	//please check before update
	if err := cr.db.Save(&class_register).Error; err!=nil {
		return err
	}

	ps.UpdateTotalCreditByID(class_register.PaymentID)

	return nil
}

func (cr *ClassRegisterService) DeleteRegisterClass(context *gin.Context) error {
	class_register := model.ClassRegister{}
	var data struct{
		ID uint `json:"id"`
	}
	ps := NewPaymentService(cr.db)

	if err := context.ShouldBindJSON(&data); err!=nil {
		return err
	}

	if result := cr.db.Delete(&class_register, data.ID); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete this register")
	}
	ps.UpdateTotalCreditByID(class_register.PaymentID)
	
	return nil
}

func (ps *PaymentService) FindClassIDWithExistsCourse(payment_id uint) uint {
	var studentID model.Payment

	if err := ps.db.Model(model.Payment{}).Where("ID = ?", payment_id).First(&studentID).Error; err != nil {
		return 0
	}

	return studentID.StudentID
}