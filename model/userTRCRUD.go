package model

import "log"

//カスタムトレーニング追加
func CreateUserTrainings(u UserTraining) error {
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

//カスタムトレーニング削除
func DeleteUserTrainings(u UserTraining) error {
	db := ConnectionByTCP()
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

func ReadUserTrainings(id string) []UserTraining {
	db := ConnectionByTCP()
	var ut []UserTraining
	_ = db.Where("user_id = ?", id).Find(&ut)
	return ut
}
