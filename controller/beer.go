package controller

import (
	"github.com/gin-gonic/gin"
	"../datasource"
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
