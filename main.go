package main

import (
	"gmo_2022_summer/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
		users.PUT("/editUser", controller.UpdateUser)
		users.POST("/getUser", controller.GetUser)
	}

	customTR := r.Group("/api/customTR")
	{
		customTR.POST("/", controller.CustomeTR)
		customTR.POST("/add", controller.AddCustomeTR)
		customTR.DELETE("/delete", controller.DeleteCustomeTR)
	}

	trainingHis := r.Group("/api/training")
	{
		trainingHis.POST("/add", controller.AddTrainingHistory)
		trainingHis.POST("/", controller.ShowTrainingHistory)
	}

	r.GET("/api/addpublictrainings/", controller.AddPublicTrainings)

	r.Run()

}
