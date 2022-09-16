package server

import (
	"github.com/gin-gonic/gin"
)

func setupBarsRouting(router *gin.Engine, app *AppConfig) {
	router.POST("/createbar", app.CreateBar)
	router.PUT("/updatebar/:id", app.UpdateBar)
	router.DELETE("/deletebar/:id", app.UpdateBar)
	router.GET("/bar/:id", app.GetBar)
}

func setupDrinksRouting(router *gin.Engine, app *AppConfig) {
	router.POST("/createdrink", app.CreateDrink)
	router.POST("/updatedrink/:id", app.UpdateDrink)
	router.GET("/drink/:id", app.GetDrink)
	router.GET("/drink/:id/ingredients", app.GetDrinkIngredients)
}
func setupIngredientsRouting(router *gin.Engine, app *AppConfig) {

}

func SetupRouting(router *gin.Engine, app *AppConfig) {
	setupBarsRouting(router, app)
	setupDrinksRouting(router, app)
	setupIngredientsRouting(router, app)
}
