package main

import (
	"gmo_2022_summer/controller"

	"github.com/gin-contrib/cors"
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
	// ここからCorsの設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})

	// cookie is not required to this endpoints
	cnr := r.Group("/api")
	{
		cnr.POST("/register", controller.Register)
		cnr.POST("/login", controller.Login)
	}

	users := r.Group("/api/users")
	{
		//users.GET("/checkDuplication/:userID", controller.CheckDuplication)
		users.PUT("/editUser", controller.UpdateUser)
		users.GET("/getUser", controller.GetUser)
	}

	customTR := r.Group("/api/customTR")
	{
		customTR.GET("/", controller.CustomeTR)
		customTR.POST("/add", controller.AddCustomeTR)
		customTR.DELETE("/delete", controller.DeleteCustomeTR)
	}

	trainingHis := r.Group("/api/training")
	{
		trainingHis.POST("/add", controller.AddTrainingHistory)
		trainingHis.GET("/", controller.ShowTrainingHistory)
	}

	// -------------------------------------------------------------------------

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
