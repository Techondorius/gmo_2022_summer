package main

import (
	"gmo_2022_summer/controller"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	// cookie is not required to this endpoints
	cnr := r.Group("/api")
	{
		cnr.POST("/register", controller.Register)
		cnr.GET("/login", controller.Login)
		//cnr.POST("/training/add", controller.TrainingAdd)
	}

	users := r.Group("/api/users")
	{
		users.GET("/checkDuplication/:userID", controller.CheckDuplication)
		users.PUT("/editUser", controller.UpdateUser)
		users.GET("/getUser", controller.GetUser)
	}

	custometr := r.Group("/api/customeTR")
	{
		//トレーニング一覧取得（後回し）
		custometr.GET("/", controller.CustomeTR)
		custometr.POST("/add", controller.AddCustomeTR)
		custometr.DELETE("/delete", controller.DeleteCustomeTR)
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
