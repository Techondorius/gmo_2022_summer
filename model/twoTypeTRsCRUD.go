package model

import "log"

// CreateUserTrainings カスタムトレーニング追加
func CreateUserTrainings(u UserTraining) error {
	db := Connection()
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ReadUserTrainings(id string) []UserTraining {
	db := Connection()
	var ut []UserTraining
	_ = db.Where("user_id = ?", id).Find(&ut)
	return ut
}

func ReadPublicTrainings() []PublicTraining {
	db := Connection()
	var pt []PublicTraining
	_ = db.Debug().Find(&pt)
	return pt
}

// DeleteUserTrainings カスタムトレーニング削除
func DeleteUserTrainings(u UserTraining) error {
	db := Connection()
	log.Println(u)
	//"10"のところにuser_id認証データ持ってくる
	result := db.Delete(&u, 3)
	log.Println(result)
	log.Println(result.Error)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
