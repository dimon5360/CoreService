package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRouting(router *gin.Engine, app *AppConfig) {

	// bars requests
	// POST requests
	router.POST("/createbar", app.CreateBar)

	// GET requests
	router.GET("/bar/:id", app.GetBar)

	// drinks requests
	// POST
	router.POST("/createdrink", app.CreateDrink)

	// GET
	router.GET("/drink/:id", app.GetDrink)
}
