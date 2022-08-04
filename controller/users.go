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
	ta := model.TrainingAddst{
		ID:       1,
		UserID:   "Pi",
		IsCustom: false,
		TLength:  60,
		TWhen:    1659592629,
	}
	cal := model.GetNameConsumptingC(ta.ID, ta.IsCustom)
	weight := model.GetUser(ta.UserID).Weight
	if ta.IsCustom {

		calorie := cal
	} else {
		c := float64(cal) * float64(weight) * float64(ta.TLength) * 1.05
		calorie := int(c / 1)
	}

	res := model.TrainingHistory{
		UserID:       ta.UserID,
		TWhen:        ta.TWhen,
		UserTraining: ta.IsCustom,
		TName:        "",
		TLength:      "",
		ConsumptingC: 0,
	}

	if err := model.CreateTrainingHistory(); err != nil {
		c.JSON(400, nil)
		return
	} else {
		c.JSON(200, nil)
	}

	c.JSON(200, gin.H{
		"Detail": map[string]any{
			"ID": 1004,
			//"Time": Time.Now(),
			"TName":   "スクワット",
			"TLength": 120,
			//"ConsumptingC": Calorie,
		}})
	c.JSON(200, gin.H{"message": "CreateTrainingHistory"})
}

func Login(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Login"})
}

func CheckDuplication(c *gin.Context) {

	c.JSON(200, gin.H{"message": true})
}

func UpdateUser(c *gin.Context) {
	type request struct {
		ID        string
		Name      string
		Birthdate int
		Sex       int
		Height    int
		Weight    int
		Objective int
		Password  string
		NPassword string
	}

	u := request{
		ID:        "Pi",
		Name:      "GHJK",
		Birthdate: 12341234,
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
	model.UpdateUser(newu)
	//log.Println(u)
	//model.CreateUser(u)
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
	td := model.ReadTrainingHistory(tt.UserID, dtstart, dtstop)
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
			"ID":           gu.ID,
			"Name":         gu.Name,
			"Birthdate":    gu.Birthdate,
			"Sex":          gu.Sex,
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
	model.CreateUser(u)
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
