package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngredientBody struct {
	Id      int32
	Title   string
	Amount  int32
	DrinkId int32
}

// / create gRPC create ingredient request from HTTP request body
// #TODO: need to implement script
func (app *AppCore) CreateIngredient(c *gin.Context) {

	c.String(http.StatusOK, "")
}

// / create gRPC update ingredient request from HTTP request body
// #TODO: need to implement script
func (app *AppCore) UpdateIngredient(c *gin.Context) {

	c.String(http.StatusOK, "")
}

// / create gRPC delete ingredient request from HTTP request body
// #TODO: need to implement script
func (app *AppCore) DeleteIngredient(c *gin.Context) {

	c.String(http.StatusOK, "")
}
