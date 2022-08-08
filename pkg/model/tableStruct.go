package model

type User struct {
	UserID        string `gorm:"primaryKey" json:"UserID"`
	Name      string `gorm:"not null" json:"Name"`
	Birthdate int    `gorm:"not null" json:"Birthdate"`
	Sex       int    `gorm:"size:4;not null" json:"Sex"`
	Height    int    `gorm:"not null" json:"Height"`
	Weight    int    `gorm:"not null" json:"Weight"`
	Objective int    `gorm:"not null" json:"Objective"`
	Password  string `gorm:"not null" json:"-"`
}

// カロリー計算方法 = Calorie * UnitTime * 回数(or n分)
type PublicTraining struct {
	TRID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Mets int    `gorm:"not null"`
}

type UserTraining struct {
	TRID      int    `gorm:"primaryKey"`
	UserID  string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Calorie int    `gorm:"not null"`
}

type TrainingHistory struct {
	TRID           int    `gorm:"primaryKey"`
	UserID       string `gorm:"not null"`
	TWhen        int    `gorm:"not null"`
	UserTraining bool   `gorm:"not null"`
	TName        string `gorm:"not null"`
	TLength      int    `gorm:"not null"`
	ConsumptingC int    `gorm:"not null"`
}

type TRLIst struct {
	TRID           int
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
	TRID       int
	UserID   string
	IsCustom bool
	TLength  int
	TWhen    int
}
