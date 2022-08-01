package main

import (
	"gmo_2022_summer/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})
	r.GET("/create", controller.CreateUser)
	r.Run()

}