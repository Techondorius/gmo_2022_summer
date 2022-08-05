package model

func CreateTrainingHistory(u TrainingHistory) error {
	db := Connection()
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// ReadTrainingHistory 始端と終端が必要
func ReadTrainingHistory(id string, start int, end int) []TrainingHistory {
	db := Connection()
	var th []TrainingHistory
	//"UO"のところは認証情報からとってくる
	_ = db.Debug().Where("user_id = ? AND t_when >= ? AND ", id, start, end).Find(&th)
	return th
}
