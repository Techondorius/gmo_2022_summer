package main

import (
	"gmo_2022_summer/controller"
	"gmo_2022_summer/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	// cookie is not required to this endpoints
	cnr := r.Group("/api")
	cnr.Use(CheckCookie())
	{
		cnr.POST("/register", controller.Register)
		cnr.GET("/login", controller.Login)
	}

	users := r.Group("/api/users")
	{
		users.GET("/checkDuplication/:userID", controller.CheckDuplication)
		users.PUT("/editUser", controller.UpdateUser)
		users.GET("/getUser", controller.GetUser)
	}

	utr := r.Group("/api/customeTR")
	utr.Use(CheckCookie())
	{
		utr.GET("/", controller.GetTRHis)
	}

	// -------------------------------------------------------------------------

	r.GET("/create", controller.CreateUser)

	r.GET("/ip", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": c.ClientIP()})
	})

	r.GET("/cookie", func(c *gin.Context) {
		a, _ := c.Cookie("Cookie")
		c.JSON(200, gin.H{"message": a})
	})

	r.GET("/gethash", func(c *gin.Context) {
		password := []byte("password")
		hashed, _ := bcrypt.GenerateFromPassword(password, 4)
		c.JSON(200, gin.H{"message": hashed})
	})

	r.GET("/decodehash", func(c *gin.Context) {
		password := []byte("password")
		hashed, _ := bcrypt.GenerateFromPassword(password, 4)
		err := bcrypt.CompareHashAndPassword(hashed, password)
		if err != nil {
			c.JSON(200, gin.H{"message": "False"})
		} else {
			c.JSON(200, gin.H{"message": "True"})
		}
	})

	r.Run()

}

func CheckCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Cookie("userID")
		token, _ := c.Cookie("token")
		if res := model.FindCookie(userID, token); res == "" {
			c.JSON(403, nil)
		}
		pw := model.GetPassword(userID)
		if err := bcrypt.CompareHashAndPassword([]byte(token), []byte(pw)); err != nil {
			c.JSON(403, nil)
			c.Abort()
		} else {

			c.Next()
		}
	}
}
