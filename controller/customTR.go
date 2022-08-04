package controller

import (
	"github.com/gin-gonic/gin"
	"gmo_2022_summer/model"
	"log"
)

//トレーニング一覧
func CustomeTR(c *gin.Context) {
	//u := model.UserTraining{
	//	UserID: "PI", //cookieから取得
	//}
	//
	//pt := model.ReadPublicTrains(123)
	//ut := model.ReadUserTrainings(u.UserID)
	//res := pt
	//asd :=
	//res = append(res, UTtoTRL(ut)...)
	//log.Println(res)

	c.JSON(200, nil)
}

func AddCustomeTR(c *gin.Context) {
	u := model.UserTraining{
		Name:    "kensui",
		Calorie: 10,
	}
	if err := c.Bind(&u); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "Update Failed"})
		return
	}
	newu := model.UserTraining{}
	newu.UserID = "PI" //cookieから取得
	newu.Name = u.Name
	newu.Calorie = u.Calorie

	model.CreateUserTrainings(newu)
	log.Println(u)
	c.JSON(200, gin.H{
		"detail": map[string]any{
			"UserId":  "PI", //Cookieから取得
			"UserTR":  true,
			"Name":    u.Name,
			"Calolie": u.Calorie,
		}})
}

func DeleteCustomeTR(c *gin.Context) {
	u := model.UserTraining{
		ID: 3,
	}
	newu := model.UserTraining{}
	newu.ID = 3
	model.DeleteUserTrainings(newu)
	log.Println(model.DeleteUserTrainings(newu))
	c.JSON(200, gin.H{
		"detail": map[string]any{
			"ID": u.ID,
		}})
}
