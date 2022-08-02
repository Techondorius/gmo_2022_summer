package model

import (
	"time"
)

type User struct{
	ID string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Sex int `gorm:"size:4;not null"`
	Height int `gorm:"not null"`
	Weight int `gorm:"not null"`
	Objective int `gorm:"not null"`
	Password string `gorm:"not null"`
}

type PublicTrainings struct{
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Calorie int `gorm:"not null"`	// 単位あたり消費カロリー
	Unit string `gorm:"not null"`	// 単位
	UnitTime int `gorm:"not null"`	// 時間単位/回数単位
}

// カロリー計算方法 = Calorie * UnitTime * 回数(or n分)