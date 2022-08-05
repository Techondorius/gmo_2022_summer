package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gmo_2022_summer/model"
	"log"
)

// AllTypeTR ConsumptingC has Mets if UserTraining is false and has Consumpting Calorie if true
type AllTypeTR struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	UserTraining bool   `json:"UserTraining"`
	ConsumptingC int    `json:"ConsumptingC"`
}

func CustomeTR(c *gin.Context) {
	type request struct {
		UserID string `json:"UserID" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		log.Println(req)
		c.JSON(400, gin.H{"Message": err})
		return
	}

	res := TrainingList(req.UserID)
	c.JSON(200, map[string]any{
		"Details": res,
	})
}

func TrainingList(userID string) []AllTypeTR {
	pt := model.ReadPublicTrainings()
	ut := model.ReadUserTrainings(userID)
	attr := []AllTypeTR{}
	for i := 0; i < len(pt); i++ {
		attr = append(attr, AllTypeTR{
			ID:           pt[i].ID,
			Name:         pt[i].Name,
			UserTraining: false,
			ConsumptingC: pt[i].Mets,
		})
	}
	for i := 0; i < len(ut); i++ {
		attr = append(attr, AllTypeTR{
			ID:           ut[i].ID,
			Name:         ut[i].Name,
			UserTraining: true,
			ConsumptingC: pt[i].Mets,
		})
	}
	return attr
}

func AddCustomeTR(c *gin.Context) {
	type request struct {
		Name         string `json:"Name" binding:"required"`
		UserID       string `json:"Userid" binding:"required"`
		ConsumptingC int    `json:"ConsumptingC" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"detail": 1})
		return
	}

	var ut model.UserTraining
	if err := copier.Copy(&ut, &req); err != nil {
		c.JSON(400, gin.H{"detail": 2})
	}

	if err := model.CreateUserTrainings(ut); err != nil {
		c.JSON(400, gin.H{"detail": 3})
		return
	} else {
		c.JSON(200, gin.H{"Detail": TrainingList(req.UserID)})
	}
}

func DeleteCustomeTR(c *gin.Context) {
	type request struct {
		ID     int    `json:"ID" binding:"required"`
		UserID string `json:"UserID" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Detail": 1})
		return
	}
	if err := model.DeleteUserTrainings(req.ID); err != nil {
		c.JSON(400, gin.H{"Detail": 2})
	} else {
		c.JSON(200, gin.H{"Detail": TrainingList(req.UserID)})
	}

}
