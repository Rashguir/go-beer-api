package controller

import (
	"beerapi/datasource"
	"beerapi/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AllBeersHandler ...
func AllBeersHandler(c *gin.Context) {
	beers := datasource.GetBeers()

	c.JSON(200, beers)
}

// GetBeerHandler ...
func GetBeerHandler(c *gin.Context) {
	beerID := c.Param("id")

	beers := datasource.GetBeers()
	// found := false
	for _, b := range beers {
		if b.Fields.ID == beerID {
			c.JSON(200, b)
			return
		}
	}

	c.JSON(404, gin.H{"status": "not found"})
}

// CountBeersHandler ...
func CountBeersHandler(c *gin.Context) {
	beers := datasource.GetBeers()

	c.JSON(200, len(beers))
}

// ImportBeerFileInDb ...
func ImportBeerFileInDb(c *gin.Context) {
	db, err := gorm.Open("postgres", "host=localhost port=54320 user=db_user dbname=db_name password=db_password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error)
		panic("Db problem")
	}
	defer db.Close()

	beers := datasource.GetBeers()

	for _, b := range beers {
		db.Create(&b.Fields)
	}

	fmt.Println("Beers imported.")

	c.JSON(201, gin.H{"status": "imported"})
}

// AllBeersDb ...
func AllBeersDb(c *gin.Context) {
	db, err := gorm.Open("postgres", "host=localhost port=54320 user=db_user dbname=db_name password=db_password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error)
		panic("Db problem")
	}
	defer db.Close()

	var beers []model.Beer
	db.Find(&beers)

	c.JSON(200, beers)
}
