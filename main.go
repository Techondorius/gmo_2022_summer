package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"time"

	"gmo_2022_summer/migration"
	"gmo_2022_summer/pkg/model"
	"gmo_2022_summer/routing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// DBマイグレーション
	// model.Connectionがエラー発生しなくなるまで=DBが立ち上がるまで待機
	// (docker composeで立ち上げると必ずdbのほうが立ち上がり遅い)
	_, dbConErr := model.Connection()
	for dbConErr != nil {
		time.Sleep(time.Second)
		_, dbConErr = model.Connection()
	}
	migration.Mig()

	r := gin.Default()

	r.Use(logger())

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// ルーティング
	routing.Routing(r)

	r.Run()
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ByteBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
		log.Println("endpoint: " + c.FullPath())
		log.Println("body" + string(ByteBody))
		c.Next()
	}
}
