package model

// "gorm.io/gorm"

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


