package model

import (

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectionByTCP() *gorm.DB {
	// dsn := "root:@tcp(db)/gin_app"
	dsn := "root:@tcp(localhost:3306)/gin_app?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func Connection() *gorm.DB {
	dsn := "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}