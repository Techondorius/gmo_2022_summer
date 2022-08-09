package model

type User struct {
	UserID       string         `gorm:"size:20;primaryKey" json:"UserID"`
	Name         string         `gorm:"not null" json:"Name"`
	Birthdate    int            `gorm:"not null" json:"Birthdate"`
	Sex          int            `gorm:"size:4;not null" json:"Sex"`
	Height       int            `gorm:"not null" json:"Height"`
	Weight       int            `gorm:"not null" json:"Weight"`
	Objective    int            `gorm:"not null" json:"Objective"`
	Password     string         `gorm:"not null" json:"-"`
	UserTraining []UserTraining `gorm:"foreignKey:UserID;references:UserID" json:"-"`
	TrainingHistory []TrainingHistory `gorm:"foreignKey:UserID;references:UserID" json:"-"`
}

// カロリー計算方法 = Calorie * UnitTime * 回数(or n分)
type PublicTraining struct {
	TRID int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Mets int    `gorm:"not null"`
}

type UserTraining struct {
	TRID    int    `gorm:"primaryKey"`
	UserID  string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Calorie int    `gorm:"not null"`
}

type TrainingHistory struct {
	TRID         int    `gorm:"primaryKey" json:"TRID"`
	UserID       string `gorm:"not null" json:"-"`
	TWhen        int    `gorm:"not null" json:"TWhen"`
	UserTraining bool   `gorm:"not null" json:"UserTraining"`
	TName        string `gorm:"not null" json:"TName"`
	TLength      int    `gorm:"not null" json:"TLength"`
	ConsumptingC int    `gorm:"not null" json:"ConsumptingC"`
}
