package model

// "gorm.io/gorm"
import(
	"log"
)

func UserCreate(u User) error {
	db := ConnectionByTCP()
	//db.Create(&product) // pass pointer of data to Create
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
func UserUpdate(u User) error {
	db := ConnectionByTCP()

	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func AddCustomeTR(u UserTraining) error {
	db := ConnectionByTCP()
	log.Println(u)
	result := db.Create(&u)
	log.Println(result)
	log.Println(result.Error)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
