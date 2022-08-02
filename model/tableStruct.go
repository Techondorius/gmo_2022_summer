package model

import (
	"time"
)

type User struct{
	ID string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Sex int `gorm:"not null"`
	Objective int `gorm:"not null"`
	Password string `gorm:"not null"`
}