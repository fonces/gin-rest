package employee

import (
	"github.com/fonces/gin-rest/models"
	"github.com/fonces/gin-rest/utils"
	"github.com/jinzhu/gorm"
)

type Employee struct {
	Id   int
	Name string
	CreatedAt string
	UpdatedAt string
}

var db *gorm.DB

func init() {
	db = model.GetInstance(&Employee{})
}

// get first row
func FindFirst(where ...interface{}) (Employee, error) {
	var employee Employee
	db.First(&employee, where...)
	return employee, db.Error
}

// get list
func Find() ([]Employee, error) {
	var employees []Employee
	db.Find(&employees)
	return employees, db.Error
}

// create row
func Create(employee *Employee) error {
	employee.CreatedAt = utils.Date.Now
	employee.UpdatedAt = utils.Date.Now
	return db.Create(&employee).Error
}

// update row
func Update(employee *Employee) error {
	employee.UpdatedAt = utils.Date.Now
	return db.Model(&employee).Update(employee).Error
}

// delete row
func Delete(id int) error {
	return db.Delete(Employee{ Id: id }).Error
}
