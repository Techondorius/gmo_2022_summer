package view

import (
	"time"

	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, m string, d any) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	log.Println("body:\n" + string(body))
	if d == nil {
		c.JSON(
			200,
			gin.H{
				"message":     m,
				"server time": time.Now().Unix(),
			},
		)
	} else {
		c.JSON(
			200,
			gin.H{
				"message":     m,
				"detail":      d,
				"server time": time.Now().Unix(),
			},
		)
	}
	return
}

func BadRequest(c *gin.Context, m string) {
	c.JSON(
		400,
		gin.H{
			"message":     m,
			"server time": time.Now().Unix(),
		},
	)
	return
}
