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

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get"})
}

func ShowTables() {
	db := model.Connection()
	rows, _ := db.Raw("show tables").Rows()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		log.Printf(table)
	}
}

func CreateUser(c *gin.Context) {
	ShowTables()
	// var u model.User
	//u := model.User
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
