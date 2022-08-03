package model

import (
	"time"
)

type User struct {
	ID              string            `gorm:"size:20;primaryKey"`
	Name            string            `gorm:"not null"`
	Birthdate       time.Time         `gorm:"not null"`
	Sex             int               `gorm:"size:4;not null"`
	Height          int               `gorm:"not null"`
	Weight          int               `gorm:"not null"`
	Objective       int               `gorm:"not null"`
	Password        string            `gorm:"not null"`
	UserTrainings   []UserTraining    `gorm:"foreignKey:UserID"`
	TrainingHistory []TrainingHistory `gorm:"foreignKey:UserID"`
}

type PublicTraining struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Mets int    `gorm:"not null"`
}

type UserTraining struct {
	ID      int    `gorm:"primaryKey"`
	UserID  string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Calorie string `gorm:"not null"`
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

type CookieData struct {
	ID     int       `gorm:"primaryKey"`
	UserID string    `gorm:"not null"`
	Cookie string    `gorm:"not null"`
	Expiry time.Time `gorm:"not null"`
}
