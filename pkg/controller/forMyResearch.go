package controller

import (
	"gmo_2022_summer/pkg/view"

	"github.com/gin-gonic/gin"
)

func AddPublicTrainings(c *gin.Context) {
	req := c.Param("int")
	view.StatusOK(c, "asdf", req)
}