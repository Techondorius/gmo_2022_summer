package controller

import (
	"github.com/gin-gonic/gin"
	"gmo_2022_summer/pkg/model"
	"log"
	"time"
)

func AddTrainingHistory(c *gin.Context) {
	type request struct {
		UserID   string `json:"UserID" binding:"required"`
		TRID       int    `json:"TRID" binding:"required"`
		IsCustom bool   `json:"IsCustom" binding:"required"`
		TLength  int    `json:"TLength" binding:"required"`
		TWhen    int    `json:"TWhen" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Detail": 1})
		return
	}
	if req.IsCustom {
		tr := model.ReadUserTrainigsByPK(req.UserID, req.TRID)
		th := model.TrainingHistory{
			UserID:       req.UserID,
			TWhen:        req.TWhen,
			UserTraining: true,
			TName:        tr.Name,
			TLength:      1,
			ConsumptingC: tr.Calorie,
		}
		if err := model.CreateTrainingHistory(th); err != nil {
			c.JSON(400, gin.H{"Detail": 2})
			return
		}
	} else {
		tr := model.ReadPublicTrainingsByPK(req.TRID)
		u := model.GetUser(req.UserID)
		th := model.TrainingHistory{
			UserID:       req.UserID,
			TWhen:        req.TWhen,
			UserTraining: false,
			TName:        tr.Name,
			TLength:      req.TLength,
			ConsumptingC: int((float64(tr.Mets) * float64(req.TLength) * float64(u.Weight) * 1.05) / 1),
		}
		if err := model.CreateTrainingHistory(th); err != nil {
			c.JSON(400, gin.H{"Detail": 2})
			return
		}
	}
	res := TodayTrainingHistory(req.UserID)
	c.JSON(200, gin.H{"Details": res})
}

func TodayTrainingHistory(userid string) []model.TrainingHistory {
	n := time.Now()
	start := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local).Unix()
	end := time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 999999, time.Local).Unix()
	return model.ReadTrainingHistory(userid, int(start), int(end))
}

func ShowTrainingHistory(c *gin.Context) {
	type request struct {
		UserID    string `json:"UserID" binding:"required"`
		StartTime int    `json:"StartTime"`
		EndTime   int    `json:"EndTime"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Detail": 1})
		return
	}
	log.Println(req.StartTime, req.EndTime)

	if req.EndTime == 0 {
		req.EndTime = 9999999999
	}

	th := model.ReadTrainingHistory(req.UserID, req.StartTime, req.EndTime)
	c.JSON(200, gin.H{"Detail": th})
}
