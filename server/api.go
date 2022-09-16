package server

import (
	"github.com/gin-gonic/gin"
)

func setupBarsRouting(router *gin.Engine, app *AppConfig) {
	router.POST("/createbar", app.CreateBar)
	router.PUT("/updatebar/:id", app.UpdateBar)
	router.DELETE("/deletebar/:id", app.DeleteBar)
	router.GET("/bar/:id", app.GetBar)
}

func setupDrinksRouting(router *gin.Engine, app *AppConfig) {
	router.POST("/createdrink", app.CreateDrink)
	router.PUT("/updatedrink/:id", app.UpdateDrink)
	router.DELETE("/deletedrink/:id", app.DeleteDrink)
	router.GET("/drink/:id", app.GetDrink)
}
func setupIngredientsRouting(router *gin.Engine, app *AppConfig) {
	router.POST("/createingredient", app.CreateIngredient)
	router.PUT("/updateingredient/:id", app.UpdateIngredient)
	router.DELETE("/deleteingredient/:id", app.DeleteIngredient)
}

func SetupRouting(router *gin.Engine, app *AppConfig) {
	setupBarsRouting(router, app)
	setupDrinksRouting(router, app)
	setupIngredientsRouting(router, app)
}
