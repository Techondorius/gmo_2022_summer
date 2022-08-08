package migration

import (
	"gmo_2022_summer/pkg/model"
	"log"
	// "gorm.io/gorm"
)

func Mig() {
	db := model.Connection()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserTraining{})
	db.AutoMigrate(&model.PublicTraining{})
	db.AutoMigrate(&model.TrainingHistory{})
	log.Println("!!")
}
