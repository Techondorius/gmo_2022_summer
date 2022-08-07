package controller

import (
	"gmo_2022_summer/model"
	"log"

	"github.com/gin-gonic/gin"
)

func ShowTables() {
	db := model.Connection()
	rows, _ := db.Raw("show tables").Rows()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		log.Printf(table)
	}
}

func CreateUser(c *gin.Context) {
	ShowTables()
	c.JSON(200, "aaa")
}
