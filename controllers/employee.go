package controllers

import (
	"gihub.com/fonces/go-db-connect/models/employee"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeController struct {}
var Employee EmployeeController

// 一覧
func (e *EmployeeController) List(c *gin.Context) {
	emps, err := employee.Find()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Employees": emps})
}

// 詳細
func (e *EmployeeController) Detail(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		return
	}

	emp, err := employee.FindFirst("id = ?", id)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Employee": emp})
}

// 作成
func (e *EmployeeController) Create(c *gin.Context) {

	emp := &employee.Employee{}
	c.Bind(emp)
	err := employee.Create(emp)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Employee": emp})
}

// 更新
func (e *EmployeeController) Update(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		return
	}

	emp := &employee.Employee{
		Id: id,
	}
	c.Bind(emp)
	err = employee.Update(emp)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Employee": emp})
}

// 削除
func (e *EmployeeController) Delete(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		return
	}

	err = employee.Delete(id)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
