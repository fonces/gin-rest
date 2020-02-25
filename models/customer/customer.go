package customer

import (
	"gihub.com/fonces/go-db-connect/models"
	"gihub.com/fonces/go-db-connect/utils"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	Id   int
	CustomerName string
	CreatedAt string
	UpdatedAt string
}

var db *gorm.DB

func init() {
	db = model.GetInstance(&Customer{})
}

// get first row
func FindFirst(where ...interface{}) (Customer, error) {
	var customer Customer
	db.First(&customer, where...)
	return customer, db.Error
}

// get list
func Find() ([]Customer, error) {
	var customers []Customer
	db.Find(&customers)
	return customers, db.Error
}

// create row
func Create(customer *Customer) error {
	customer.CreatedAt = utils.Date.Now
	customer.UpdatedAt = utils.Date.Now
	return db.Create(&customer).Error
}

// update row
func Update(customer *Customer) error {
	customer.UpdatedAt = utils.Date.Now
	return db.Model(&customer).Update(customer).Error
}

// delete row
func Delete(id int) error {
	return db.Delete(Customer{ Id: id }).Error
}
