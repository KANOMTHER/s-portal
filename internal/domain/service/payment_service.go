package service

import (
	// "fmt"

	// "github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"s-portal/internal/domain/model"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{
		db: db,
	}
}

func (ps *PaymentService) GetAndCreatePaymentByID(payment model.Payment) (model.Payment, error) {
	empty := model.Payment{}

	if err := ps.db.Where(payment).FirstOrCreate(&payment).Error; err!=nil {
		return empty, err
	}

	return payment, nil
}

func (ps *PaymentService) UpdateTotalCreditByID(payment_id uint) (error) {
	//input = payment_id

	//process
	/*
	 * get all class_id from class_register with payment_id
	 * get all course_id from class with class_id
	 * get all credit from course with course_id
	 */

	var class_id []uint
	if err := ps.db.Model(model.ClassRegister{}).Where("payment_id = ?", payment_id).Select("class_id").Find(&class_id).Error; err!=nil {
		return err
	}

	var course_id []uint32
	if err := ps.db.Model(model.Class{}).Where("ID in ?", class_id).Select("course_id").Find(&course_id).Error; err!=nil {
		return err
	}

	var credit []float64
	if err := ps.db.Model(model.Course{}).Where("ID in ?", course_id).Select("credit").Find(&credit).Error; err!=nil {
		return err
	}

	sum_credit := 0.0
	for i := 0; i < len(credit); i++ {
		sum_credit += credit[i]
	}
	if err := ps.db.Model(model.Payment{}).Where("ID = ?", payment_id).Update("credit = ?", sum_credit).Error; err!=nil {
		return err
	}

	return nil
}

func (ps *PaymentService) UpdatePriceByID(payment_id uint) (error) {

	return nil
}