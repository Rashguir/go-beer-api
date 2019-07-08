package main

import (
	"./controller"
	"./security"
	"github.com/gin-gonic/gin"
	"./model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

var routes = []route{
	route{"AllBeers", "GET", "/beers", controller.AllBeersHandler},
	route{"CountBeers", "GET", "/beersCount", controller.CountBeersHandler},
	route{"GetBeer", "GET", "/beers/:id", controller.GetBeerHandler},
	route{"AllBeersDb", "GET", "/db/beers", controller.AllBeersDb},
	route{"ImportBeerFileInDb", "GET", "/import", controller.ImportBeerFileInDb},
}

func main() {
	model.InitialMigration()

	router := gin.Default()

	router.GET("/login", gin.BasicAuth(gin.Accounts{"toto": "qwerty"}), security.GenerateToken)
	router.HEAD("/verify", security.VerifyToken)

	for _, r := range routes {
		router.Handle(r.Method, r.Pattern, security.VerifyToken, r.HandlerFunc)
	}
	router.Run(":8080")
}
