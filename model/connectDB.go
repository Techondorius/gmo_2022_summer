package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionByTCP() *gorm.DB {
	// dsn := "root:@tcp(db)/gin_app"
	dsn := "root:asdl0606@tcp(163.44.255.164:3306)/gin_app?charset=utf8&parseTime=True&loc=Local"
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
