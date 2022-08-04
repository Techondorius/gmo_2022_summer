package model

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Sex       int       `gorm:"size:4;not null"`
	Height    int       `gorm:"not null"`
	Weight    int       `gorm:"not null"`
	Objective int       `gorm:"not null"`
	Password  string    `gorm:"not null"`
}

type UpdateUser struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Sex       int       `gorm:"size:4;not null"`
	Height    int       `gorm:"not null"`
	Weight    int       `gorm:"not null"`
	Objective int       `gorm:"not null"`
	Password  string
	NPassword string
}

// カロリー計算方法 = Calorie * UnitTime * 回数(or n分)
type PublicTraining struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Mets int    `gorm:"not null"`
}

type UserTraining struct {
	ID      int    `gorm:"primaryKey"`
	UserID  string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Calorie int    `gorm:"not null"`
}

type TrainingHistory struct {
	ID           int       `gorm:"primaryKey"`
	UserID       string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UserTraining bool      `gorm:"not null"`
	TName        string    `gorm:"not null"`
	TLength      string    `gorm:"not null"`
	ConsumptingC int       `gorm:"not null"`
}

type TRLIst struct {
	ID           int
	Name         string
	UserTR       bool
	ConsumptingC int
}

type TrainingTime struct {
	UserID    string
	StartTime int
	EndTime   int
}

type TrainingAddst struct {
	ID int
	IsCustom bool
	TLength int
	TWhen int
}
