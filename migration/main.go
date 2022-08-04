package main

import (
	"gmo_2022_summer/model"
	"log"
	// "gorm.io/gorm"
)

func main() {
	db := model.ConnectionByTCP()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UpdateUser{})
	db.AutoMigrate(&model.UserTraining{})
	db.AutoMigrate(&model.PublicTraining{})
	db.AutoMigrate(&model.TrainingHistory{})
	log.Println("!!")
}
