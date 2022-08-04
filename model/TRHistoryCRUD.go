package model

import "time"

//トップページの履歴追加
func CreateTrainingHistory(u TrainingHistory) error {
	db := ConnectionByTCP()
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//user_idとtraining_dateを指定してTrainingHistoryから情報を抜き出す
func ReadTrainingHistory(id string, start time.Time, stop time.Time) []TrainingHistory {
	db := ConnectionByTCP()
	var th []TrainingHistory
	//"UO"のところは認証情報からとってくる
	_ = db.Debug().Where("user_id = ? AND ? <= created_at <= ?", id, start, stop).Find(&th)
	return th
}
