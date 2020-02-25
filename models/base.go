package model

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
  USER = "docker"
  PASSWORD = "docker"
  DATABASE = "example"
)

var instance *gorm.DB

func init() {
  db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s", USER, PASSWORD, DATABASE))
  if err != nil {
    panic(err)
  }
  instance = db
}

func GetInstance(values ...interface{}) *gorm.DB {
  return instance.AutoMigrate(values...)
}
