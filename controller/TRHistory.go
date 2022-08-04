package controller

import (
	"github.com/gin-gonic/gin"
	"gmo_2022_summer/model"
	"log"
)

func GetTRHis(c *gin.Context) {
	type request struct {
		ID        int `json:"ID"`
		StartTime int `json:"StartTime"`
		EndTime   int `json:"EndTime"`
	}
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}

	res, err := model.ReadHistoryByUser(req.ID)
	if err != nil {
		c.JSON(400, nil)
	} else {
		log.Println(res)
		c.JSON(200, res)
	}
}

func CreateTRHis(c *gin.Context) {

}

func UpdateTRHis(c *gin.Context) {

}

func DeleteTRHis(c *gin.Context) {

}
