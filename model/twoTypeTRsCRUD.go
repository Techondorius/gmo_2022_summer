package model

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
func DeleteUserTrainings(u int) error {
	db := Connection()
	result := db.Delete(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
