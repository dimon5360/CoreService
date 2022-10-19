package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupBarsRouting(router *gin.Engine, app *AppCore) {
	router.GET("/createbar", app.CreateBarForm)
	router.POST("/createbar", app.CreateBar)
	router.PUT("/updatebar/:id", app.UpdateBar)
	router.DELETE("/deletebar/:id", app.DeleteBar)
	router.GET("/bar/:id", app.GetBar)
}

func setupDrinksRouting(router *gin.Engine, app *AppCore) {
	router.POST("/createdrink", app.CreateDrink)
	router.PUT("/updatedrink/:id", app.UpdateDrink)
	router.DELETE("/deletedrink/:id", app.DeleteDrink)
	router.GET("/drink/:id", app.GetDrink)
}
func setupIngredientsRouting(router *gin.Engine, app *AppCore) {
	router.POST("/createingredient", app.CreateIngredient)
	router.PUT("/updateingredient/:id", app.UpdateIngredient)
	router.DELETE("/deleteingredient/:id", app.DeleteIngredient)
}

func SetupRouting(router *gin.Engine, app *AppCore) {
	router.GET("/", func(c *gin.Context) {
		var testText = "test request"
		app.Logger.Write(fmt.Sprintf("test log record: %s", testText))
		c.String(http.StatusOK, testText)
	})

	setupBarsRouting(router, app)
	setupDrinksRouting(router, app)
	setupIngredientsRouting(router, app)
}

func (app *AppCore) InitRouter() {

	router := gin.Default()

	router.SetTrustedProxies([]string{"localhost"})
	router.LoadHTMLGlob("static/*.html")

	SetupRouting(router, app)
	app.Router = router
}
