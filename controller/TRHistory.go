package controller

import (
	"github.com/gin-gonic/gin"
	"gmo_2022_summer/model"
)

func GetTRHis(c *gin.Context) {
	id, _ := c.Get("ID")
	res, err := model.ReadHistoryByUser(id)
	if err != nil {
		c.Set("detail", nil)
		c.Set("err", true)
	} else {
		c.Set("detail", res)
		c.Set("err", false)
	}
}

func CreateTRHis(c *gin.Context) {

}

func UpdateTRHis(c *gin.Context) {

}

func DeleteTRHis(c *gin.Context) {

}
