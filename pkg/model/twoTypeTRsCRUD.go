package model

import "log"

// CreateUserTrainings カスタムトレーニング追加
func CreateUserTrainings(u UserTraining) error {
	db, _ := Connection()
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func CreatePublicTrainings(p PublicTraining) error {
	db, _ := Connection()
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ReadUserTrainings(userid string) []UserTraining {
	db, _ := Connection()
	var ut []UserTraining
	_ = db.Where("user_id = ?", userid).Find(&ut)
	return ut
}

func ReadUserTrainigsByPK(userid string, pk int) UserTraining {
	db, _ := Connection()
	var ut UserTraining
	result := db.Debug().Where("id = ? AND user_id = ?", pk, userid).Find(&ut)
	log.Println(result)
	if result.Error != nil {
		return UserTraining{}
	} else {
		return ut
	}
}

func ReadPublicTrainings() []PublicTraining {
	db, _ := Connection()
	var pt []PublicTraining
	_ = db.Debug().Find(&pt)
	return pt
}

func ReadPublicTrainingsByPK(id int) PublicTraining {
	db, _ := Connection()
	var pt PublicTraining
	_ = db.Debug().Where("id = ?", id).Find(&pt)
	return pt
}

// DeleteUserTrainings カスタムトレーニング削除
func DeleteUserTrainings(u int) error {
	db, _ := Connection()
	result := db.Delete(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
