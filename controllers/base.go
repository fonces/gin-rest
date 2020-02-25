package controllers

import (
  "github.com/gin-gonic/gin"
  "strconv"
)

// URLからidを取得する
func getId(c *gin.Context) (int, error) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
  }
  return id, err
}
