package controllers

import (
  "github.com/fonces/gin-rest/models/customer"
  "github.com/gin-gonic/gin"
  "net/http"
)

type CustomerController struct {}
var Customer CustomerController

// 一覧
func (e *CustomerController) List(c *gin.Context) {
  csts, err := customer.Find()
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }
  c.JSON(http.StatusOK, gin.H{"Customers": csts})
}

// 詳細
func (e *CustomerController) Detail(c *gin.Context) {
  id, err := getId(c)
  if err != nil {
    return
  }

  cst, err := customer.FindFirst("id = ?", id)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }
  c.JSON(http.StatusOK, gin.H{"Customer": cst})
}

// 作成
func (e *CustomerController) Create(c *gin.Context) {
  cst := &customer.Customer{}
  c.Bind(cst)
  err := customer.Create(cst)

  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }
  c.JSON(http.StatusOK, gin.H{"Customer": cst})
}

// 更新
func (e *CustomerController) Update(c *gin.Context) {
  id, err := getId(c)
  if err != nil {
    return
  }

  cst := &customer.Customer{
    Id: id,
  }
  c.Bind(cst)
  err = customer.Update(cst)

  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }
  c.JSON(http.StatusOK, gin.H{"Customer": cst})
}

// 削除
func (e *CustomerController) Delete(c *gin.Context) {
  id, err := getId(c)
  if err != nil {
    return
  }

  err = customer.Delete(id)

  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }
  c.JSON(http.StatusOK, gin.H{})
}
