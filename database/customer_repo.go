package database

import (
	"fmt"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	db.AutoMigrate(&Customer{}, &ElectricBill{})
	return &CustomerRepo{
		db: db,
	}
}
func (c CustomerRepo) CreateCustomer(model Customer) error {
	err := c.db.Create(&model).Error
	if err != nil {
		fmt.Println("fail%v", err)
		return err
	}
	fmt.Println("thanh cong")
	return nil
}

func (c CustomerRepo) DeleteCustomer(ID int) error {
	err := c.db.Where("id_meter=?", ID).Delete(&Customer{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepo) UpdateCustomer(model Customer) error {
	err := c.db.Model(&model).Omit("id_meter").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
func (c CustomerRepo) GetListCustomer() ([]Customer, error) {
	var models []Customer
	err := c.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (c CustomerRepo) GetCustomer(id int64) (*Customer, error) {
	var model Customer
	err := c.db.Where("id_meter= ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (b ElectricBill) Amount() int64 {
	amou := 5 * (b.NewNumber - b.OldNumber)
	return amou

}
