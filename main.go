package main

import (
	"./controller"
	"./security"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

var routes = []route{
	route{"AllBeers", "GET", "/beers", controller.AllBeersHandler},
	route{"GetBeer", "GET", "/beers/:id", controller.GetBeerHandler},
}

func main() {
	router := gin.Default()

	router.GET("/login", gin.BasicAuth(gin.Accounts{"toto": "qwerty"}), security.GenerateToken)

	for _, r := range routes {
		router.Handle(r.Method, r.Pattern, r.HandlerFunc)
	}
	router.Run(":8080")
}
