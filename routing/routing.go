package routing

import (
	"github.com/gin-gonic/gin"
	"gmo_2022_summer/pkg/controller"
	"log"
)

func Routing(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		log.Println("name: " + c.Query("name"))
		log.Println("age: " + c.Query("age"))
		c.JSON(200, gin.H{
			"message": "good",
			"name":    c.Query("name"),
			"age":     c.Query("age"),
		})
	})

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
}
