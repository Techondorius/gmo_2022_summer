package view

import (
	"time"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, m string, d any) {
	if d == nil {
		Logger(m)
		c.JSON(
			200,
			gin.H{
				"message":     m,
				"server_time": time.Now().Unix(),
			},
		)
	} else {
		Logger(m)
		c.JSON(
			200,
			gin.H{
				"message":     m,
				"detail":      d,
				"server_time": time.Now().Unix(),
			},
		)
	}
	return
}

func BadRequest(c *gin.Context, m string) {
	Logger(m)
	c.JSON(
		400,
		gin.H{
			"message":     m,
			"server_time": time.Now().Unix(),
		},
	)
	return
}
