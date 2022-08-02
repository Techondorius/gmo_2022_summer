package main

import (
	"log"
	"gmo_2022_summer/model"

	"gorm.io/gorm"
)

func main(){
	db := model.ConnectionByTCP()
	db.AutoMigrate(&model.User{})
	log.Println("!!")
}

type sampleeee struct{
	gorm.Model
	Name string
}