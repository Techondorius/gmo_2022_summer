package model

import (
	"errors"
)

func CreateHistory(history TrainingHistory) error {
	db := Connection()
	result := db.Create(history)
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	} else {
		return nil
	}
}

func ReadHistoryByUser(userID any) ([]TrainingHistory, error) {
	db := Connection()
	var h []TrainingHistory
	result := db.Where("userID = ?", userID).Find(&h)
	if err := result.Error; err != nil {
		return nil, err
	} else {
		return h, nil
	}
}

func UpdateHistoryByID(hid int) error {
	//db := Connection()
	return nil
}
