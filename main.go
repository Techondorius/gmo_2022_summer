package main

import (
	"gmo_2022_summer/controller"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})
	r.GET("/create", controller.CreateUser)
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": c.ClientIP()})
	})
	r.GET("/cookie", func(c *gin.Context) {
		a, _ := c.Cookie("Cookie")
		c.JSON(200, gin.H{"message":a})
	})
	r.GET("/gethash", func(c *gin.Context) {
		password := []byte("password")
		hashed, _ := bcrypt.GenerateFromPassword(password, 4)
		c.JSON(200, gin.H{"message":hashed})
	})
	r.GET("/decodehash", func(c *gin.Context) {
		password := []byte("password")
		hashed, _ := bcrypt.GenerateFromPassword(password, 4)
		err := bcrypt.CompareHashAndPassword(hashed, password)
		if err != nil {
			c.JSON(200, gin.H{"message":"False"})
		} else {
			c.JSON(200, gin.H{"message":"True"})
		}
	})
	r.Run()

}