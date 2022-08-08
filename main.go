package main

import (
	// "time"
	"log"

	"gmo_2022_summer/pkg/controller"
	"gmo_2022_summer/migration"
	"gmo_2022_summer/pkg/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// for i := 0; i < 20; i ++{
	// 	log.Println(i)
	// 	time.Sleep(time.Second)
	// }
	db := model.Connection()
	res := db.Exec("SHOW TABLES")
	log.Println(res)
	if res.RowsAffected == 0 {
		migration.Mig()
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(logger)

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

	r.GET("/api/asdf/:int", controller.AddPublicTrainings)
	r.Run()

}

func logger(c *gin.Context) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])
	log.Println("Request full path: " + c.Request.Host + c.Request.URL.Path)
	log.Println("Request body: " + b)
	c.Next()
}
