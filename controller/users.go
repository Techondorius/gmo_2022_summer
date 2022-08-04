package controller

import (
	"gmo_2022_summer/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Register(c *gin.Context) {
	type request struct {
		ID        string `json:"ID"`
		Name      string `json:"Name"`
		Birthdate int    `json:"Birthdate"`
		Sex       int    `json:"Sex"`
		Height    int    `json:"Height"`
		Weight    int    `json:"Weight"`
		Password  string `json:"Password"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}
	var u model.User
	if err := copier.Copy(&u, &req); err != nil {
		c.JSON(400, nil)
		return
	}

	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 4)
	u.Password = string(hashedPW)

	// 年齢計算
	bDtimeTime := time.Unix(int64(u.Birthdate), 0)
	age := RoundTime((time.Now()).Sub(bDtimeTime).Seconds() / 31207680)

	// objective計算
	if true {
		if u.Sex == 1 {
			if age <= 7 {
				u.Objective = 1550
			} else if age <= 9 {
				u.Objective = 1850
			} else if age <= 11 {
				u.Objective = 2250
			} else if age <= 14 {
				u.Objective = 2600
			} else if age <= 17 {
				u.Objective = 2800
			} else if age <= 29 {
				u.Objective = 2650
			} else if age <= 49 {
				u.Objective = 2700
			} else if age <= 64 {
				u.Objective = 2600
			} else if age <= 74 {
				u.Objective = 2400
			} else {
				u.Objective = 2100
			}
		} else {
			if age <= 7 {
				u.Objective = 1450
			} else if age <= 9 {
				u.Objective = 1700
			} else if age <= 11 {
				u.Objective = 2100
			} else if age <= 14 {
				u.Objective = 2400
			} else if age <= 17 {
				u.Objective = 2300
			} else if age <= 29 {
				u.Objective = 2000
			} else if age <= 49 {
				u.Objective = 2050
			} else if age <= 64 {
				u.Objective = 1950
			} else if age <= 74 {
				u.Objective = 1850
			} else {
				u.Objective = 1650
			}
		}
	}

	if err := model.CreateUser(u); err != nil {
		c.JSON(400, gin.H{"message": "ID might be already taken"})
		return
	}

	c.JSON(200, u)
}

//トップページのトレーニング登録画面
/*
//func TrainingAdd(c *gin.Context) {
//	ta := model.TrainingAddst{
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
*/

func TrainingAdd(c *gin.Context) {
	return
}

func Login(c *gin.Context) {
	type request struct {
		ID       string `json:"ID"`
		Password string `json:"Password"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}
	u := model.GetUser(req.ID)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(403, gin.H{
			"Message": "Password is wrong",
		})
		return
	}
	c.JSON(200, gin.H{"Detail": true})
}

func CheckDuplication(c *gin.Context) {

	c.JSON(200, gin.H{"message": true})
}

func RoundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func UpdateUser(c *gin.Context) {

	type request struct {
		ID        string `json:"ID"`
		Name      string `json:"Name"`
		Birthdate int    `json:"Birthdate"`
		Sex       int    `json:"Sex"`
		Height    int    `json:"Height"`
		Weight    int    `json:"Weight"`
		Objective int    `json:"Objective"`
		Password  string `json:"Password"`
		NPassword string `json:"NPassword"`
	}
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, nil)
		return
	}

	u := req

	newu := model.User{}
	newu.ID = u.ID
	newu.Name = u.Name
	newu.Birthdate = u.Birthdate
	newu.Sex = u.Sex
	newu.Height = u.Height
	newu.Weight = u.Weight
	newu.Objective = u.Objective
	newu.Password = u.Password

	/*
		// 年齢計算
		bDtimeTime := time.Unix(int64(newu.Birthdate), 0)
		age := RoundTime((time.Now()).Sub(bDtimeTime).Seconds() / 31207680)

		// objective計算
		if u.Sex == 1 {
			if age <= 7 {
				newu.Objective = 1550
			} else if age <= 9 {
				newu.Objective = 1850
			} else if age <= 11 {
				newu.Objective = 2250
			} else if age <= 14 {
				newu.Objective = 2600
			} else if age <= 17 {
				newu.Objective = 2800
			} else if age <= 29 {
				newu.Objective = 2650
			} else if age <= 49 {
				newu.Objective = 2700
			} else if age <= 64 {
				newu.Objective = 2600
			} else if age <= 74 {
				newu.Objective = 2400
			} else {
				newu.Objective = 2100
			}
		} else {
			if age <= 7 {
				newu.Objective = 1450
			} else if age <= 9 {
				newu.Objective = 1700
			} else if age <= 11 {
				newu.Objective = 2100
			} else if age <= 14 {
				newu.Objective = 2400
			} else if age <= 17 {
				newu.Objective = 2300
			} else if age <= 29 {
				newu.Objective = 2000
			} else if age <= 49 {
				newu.Objective = 2050
			} else if age <= 64 {
				newu.Objective = 1950
			} else if age <= 74 {
				newu.Objective = 1850
			} else {
				newu.Objective = 1650
			}
		}
	*/

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
		ID:   "Pi",
		Name: "ASDF",
		//Birthdate: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		Sex:      1,
		Height:   169,
		Weight:   55,
		Password: "Raspberry",
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

type TRLIst struct {
	ID           int
	Name         string
	UserTR       bool
	ConsumptingC int
}

/*
func PTtoTRL(pt model.PublicTraining) []TRLIst {
	var trl []TRLIst
	log.Println(pt)
	tr := TRLIst{
		ID:           pt.ID,
		Name:         pt.Name,
		UserTR:       false,
		ConsumptingC: pt.Mets,
	}
	}
	return trl
}*/

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
