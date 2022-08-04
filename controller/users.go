package controller

import (
	"gmo_2022_summer/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Register"})
}

//トップページのトレーニング登録画面

func TrainingAdd(c *gin.Context) {
	u := model.TrainingAddst{
		ID:       1,
		IsCustom: false,
		TLength:  60,
		TWhen:    1659592629,
	}
	model.GetNameConsumptingC(u.ID, u.IsCustom)
	/*
		newu := model.TrainingHistory{}

		newu.UserID = "PI"
		newu.CreatedAt = time.Now()
		newu.UserTraining = u.IsCustom
		newu.TName = //name
		newu.TLength = u.TLength
		newu.ConsumptingC = */
	if err := c.Bind(&u); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "Update Failed"})
		return
	}
	//db := model.ConnectionByTCP()
	//今日の

	//var i model.UserTraining
	//result := db.Where("user_id = ?", "1").Find(&i)
	//cal := result["calorie"]

	log.Println(u)
	c.JSON(200, gin.H{
		"Detail": map[string]any{
			"ID":          1,
			"Name":        "Pi",
			"Birthdate":   2002 - 1 - 1,
			"Sex":         1,
			"ConsumptedC": 500,
		}})
	c.JSON(200, gin.H{"message": "TrainingAdd"})
}

func Login(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Login"})
}

func CheckDuplication(c *gin.Context) {

	c.JSON(200, gin.H{"message": true})
}

func UpdateUser(c *gin.Context) {
	u := model.UpdateUser{
		ID:        "Pi",
		Name:      "GHJK",
		Birthdate: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		Sex:       1,
		Height:    169,
		Weight:    55,
		Objective: 100,
		Password:  "Raspberry",
		NPassword: "R4spberry",
	}

	newu := model.User{}
	newu.ID = u.ID
	newu.Name = u.Name
	newu.Birthdate = u.Birthdate
	newu.Sex = u.Sex
	newu.Height = u.Height
	newu.Weight = u.Weight
	newu.Objective = u.Objective
	newu.Password = u.Password

	//db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//log.Println(u.ID)

	if err := c.Bind(&u); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "Register Failed"})
		return
	}
	model.UserUpdate(newu)
	//log.Println(u)
	//model.UserCreate(u)
	log.Println(u)
	c.JSON(200, gin.H{
		"detail": map[string]any{
			"ID":        u.ID,
			"detail":    u.Name,
			"BirthDate": u.Birthdate,
			"Sex":       u.Sex,
			"Height":    u.Height,
			"Weight":    u.Weight,
			"Password":  u.Password,
			"NPassword": u.NPassword,
		}})
	c.JSON(200, gin.H{"message": "Update"})
}

//user_idの取り方
//今日のカロリーを算出する
func GetUser(c *gin.Context) {
	tt := model.TrainingTime{
		UserID:    "UO",
		StartTime: 1659512053,
		EndTime:   1659684853,
	}
	u := model.User{}
	gu := model.GetUser(tt.UserID)
	//今日のカロリーを取得したい
	dtstart := time.Unix(int64(tt.StartTime), 0)
	dtstop := time.Unix(int64(tt.EndTime), 0)
	td := model.PeriodData(tt.UserID, dtstart, dtstop)
	calorie := 0
	log.Println(td)
	for i := 0; i < len(td); i++ {
		log.Println(td[i])
		calorie += td[i].ConsumptingC
	}
	log.Println(calorie)
	// type gu is []User
	if err := c.Bind(&u); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "Update Failed"})
		return
	}
	c.JSON(200, gin.H{
		"Detail": map[string]any{
			"ID":           gu[0].ID,
			"Name":         gu[0].Name,
			"Birthdate":    gu[0].Birthdate,
			"Sex":          gu[0].Sex,
			"Consumpted_C": calorie,
		}})
}

func CreateUser(c *gin.Context) {
	u := model.User{
		ID:        "Pi",
		Name:      "ASDF",
		Birthdate: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		Sex:       1,
		Height:    169,
		Weight:    55,
		Password:  "Raspberry",
	}
	if err := c.Bind(&u); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "Update Failed"})
		return
	}
	//db.Create(&product) // pass pointer of data to Create
	model.UserCreate(u)
	log.Println(u)
	c.JSON(200, gin.H{
		"detail": map[string]any{
			"ID":        u.ID,
			"detail":    u.Name,
			"BirthDate": u.Birthdate,
			"Sex":       u.Sex,
			"Height":    u.Height,
			"Weight":    u.Objective,
		}})
}

//トレーニング一覧
func CustomeTR(c *gin.Context) {
	u := model.UserTraining{
		UserID: "PI", //cookieから取得
	}

	pt := model.ReadPublicTrainigs()
	ut := model.ReadUserTrainings(u.UserID)
	res := PTtoTRL(pt)
	res = append(res, UTtoTRL(ut)...)
	log.Println(res)

	c.JSON(200, res)
}

type TRLIst struct {
	ID           int
	Name         string
	UserTR       bool
	ConsumptingC int
}

func PTtoTRL(pt []model.PublicTraining) []TRLIst {
	var trl []TRLIst
	log.Println(len(pt))
	for i := 0; i < len(pt); i++ {
		tr := TRLIst{
			ID:           pt[i].ID,
			Name:         pt[i].Name,
			UserTR:       false,
			ConsumptingC: pt[i].Mets,
		}
		trl = append(trl, tr)
	}
	return trl
}

func UTtoTRL(ut []model.UserTraining) []TRLIst {
	var trl []TRLIst
	log.Println(len(ut))
	for i := 0; i < len(ut); i++ {
		log.Println(ut[i])
		tr := TRLIst{
			ID:           ut[i].ID,
			Name:         ut[i].Name,
			UserTR:       true,
			ConsumptingC: ut[i].Calorie,
		}
		trl = append(trl, tr)
	}
	return trl
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

	model.AddCustomeTR(newu)
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
	model.DeleteCustomeTR(newu)
	log.Println(model.DeleteCustomeTR(newu))
	c.JSON(200, gin.H{
		"detail": map[string]any{
			"ID": u.ID,
		}})
}
