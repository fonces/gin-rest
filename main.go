package main

import (
  "github.com/fonces/gin-rest/controllers"
  "github.com/fonces/gin-rest/utils"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
)

func main() {
	r := gin.Default()
	r.Use(eventEnhanceMiddleware())

	// employee route
	r.GET("/employee", controllers.Employee.List)
	r.GET("/employee/:id", controllers.Employee.Delete)
	r.PUT("/employee", controllers.Employee.Create)
	r.PATCH("/employee/:id", controllers.Employee.Update)
  r.DELETE("/employee/:id", controllers.Employee.Delete)

	// customer route
	r.GET("/customer", controllers.Customer.List)
	r.GET("/customer/:id", controllers.Customer.Delete)
	r.PUT("/customer", controllers.Customer.Create)
	r.PATCH("/customer/:id", controllers.Customer.Update)
	r.DELETE("/customer/:id", controllers.Customer.Delete)

  // 0.0.0.0:8080
	r.Run()
}

func eventEnhanceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// utils初期化
		utils.Date.EventInit()

		// execute handler
		c.Next()

		// error handling
		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
}
