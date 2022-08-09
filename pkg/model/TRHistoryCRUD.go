package model

func CreateTrainingHistory(u TrainingHistory) error {
	db, _ := Connection()
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// ReadTrainingHistory 始端と終端が必要
func ReadTrainingHistory(trid string, start int, end int) []TrainingHistory {
	db, _ := Connection()
	var th []TrainingHistory
	//"UO"のところは認証情報からとってくる
	_ = db.Debug().Where("user_id = ? AND t_when >= ? AND ", trid, start, end).Find(&th)
	return th
}
