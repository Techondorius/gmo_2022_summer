package model

import (
	"time"
)

func FindCookie(id string, cookie string) string {
	db := Connection()
	var c CookieData
	result := db.Where("user_id = ? AND Cookie = ? AND expiry > ?", id, cookie, time.Now()).Find(&c)
	if result.RowsAffected != 1 {
		return ""
	} else {
		return c.UserID
	}
}

func GetPassword(id string) string {
	db := Connection()
	var c User
	result := db.Where("id = ?", id).Find(&c)
	if err := result.Error; err != nil {
		return ""
	} else {
		return c.Password
	}
}
