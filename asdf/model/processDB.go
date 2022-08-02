package model

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func UserRegister(){
	db := Connection()
	//product := Product{Code: "paper", Price: 80}
	r := gin.Default()
    r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "good"})
	})
	//db.Create(&product) // pass pointer of data to Create
	db.Select("code", "price").Create(&product)
	//result := db.First(&product)
	//println(result)

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	/*
		db.Create(&Product{Code: "D42", Price: 100})
		println("created")
	*/
	// Read

	//var product Product
	/*
		db.First(&product, 1)                 // find product with integer primary key
		db.First(&product, "code = ?", "D42") // find product with code D42

			// Update - update product's price to 200
			db.Model(&product).Update("Price", 200)
			// Update - update multiple fields
			db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
			db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
			println("updated")

		// Delete - delete product
		db.Delete(&product, 1)
		println("deleted")
	*/
}