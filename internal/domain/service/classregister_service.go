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

	if err := context.ShouldBindJSON(&payment); err!=nil {
		return nil, err
	}

	//replace with GetPaymentByID in payment
	if err := cr.db.Where(payment).First(&payment).Error; err!=nil {
		return nil, err
	}

	var result []model.ClassRegister

	if err := cr.db.Where("payment_id = ?", payment.ID).Find(&result).Error; err!=nil {
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

	if err := context.ShouldBindJSON(&data); err!=nil {
		return err
	}

	if data.ClassID == 0 || data.Semester == 0 || data.StudentID == 0 || data.Year == 0 {
		return fmt.Errorf("invalid input")
	}

	//replace with GetPaymentByID in payment
	payment := model.Payment{
		StudentID: data.StudentID,
		Semester: int(data.Semester),
		Year: int(data.Year),
	}
	if err := cr.db.Where(payment).FirstOrCreate(&payment).Error; err!=nil {
		return err
	}

	//require step for initial class_grade
	// -> insert after end phase, or insert after create (this method need to think about update and delete)

	class_register := model.ClassRegister{
		PaymentID: payment.ID,
		ClassID: data.ClassID,
	}
	if err := cr.db.Where(class_register).FirstOrCreate(&class_register).Error ; err!=nil {
		return err
	}

	return nil
}

func (cr *ClassRegisterService) UpdateRegisterClass(context *gin.Context) error { 
	id := context.Param("id")

	class_register := model.ClassRegister{}

	if err := cr.db.First(&class_register, id).Error; err!=nil {
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

	if err := cr.db.Save(&class_register).Error; err!=nil {
		return err
	}

	return nil
}

func (cr *ClassRegisterService) DeleteRegisterClass(context *gin.Context) error {
	class_register := model.ClassRegister{}

	id := context.Param("id")

	if result := cr.db.Delete(&class_register, id); result.RowsAffected < 1 {
		return fmt.Errorf("were not able to delete this register")
	}
	
	return nil
}