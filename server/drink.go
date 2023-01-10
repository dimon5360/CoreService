package server

import (
	"app/main/postgres"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DrinkBody struct {
	Title       string
	Price       string
	DrinkType   uint32
	Description string
	BarId       string
	Ingredients []IngredientBody
}

// GET request - form to create POST request for new bar
func (app *AppCore) CreateDrinkForm(c *gin.Context) {

	c.HTML(http.StatusOK, "create_drink.html", gin.H{
		"content": "This is an index page...",
	})
}

// gRPC request to create drink from HTTP request body
func (app *AppCore) CreateDrink(c *gin.Context) {

	body := DrinkBody{}

	b, _ := io.ReadAll(c.Request.Body)
	if err := json.Unmarshal(b, &body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request = &postgres.CreateDrinkRequest{
		Title:       body.Title,
		Price:       body.Price,
		DrinkType:   postgres.DrinkType(body.DrinkType),
		Description: body.Description,
		BarId:       body.BarId,
		Ingredients: []*postgres.CreateIngredientRequest{},
	}

	for _, ingredient := range body.Ingredients {
		var tmp = &postgres.CreateIngredientRequest{
			Title:  ingredient.Title,
			Amount: fmt.Sprintf("%d", ingredient.Amount),
		}
		request.Ingredients = append(request.Ingredients, tmp)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.CreateDrink(ctx, request)

	if err != nil {
		log.Fatalf("could not create drink: %v", err)
		c.String(http.StatusInternalServerError, "Creating drink failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusCreated, r.String())
}

// / create gRPC update drink request from HTTP request body
// / established, remove comment later
func (app *AppCore) UpdateDrink(c *gin.Context) {

	type Req struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req Req
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	body := DrinkBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request = &postgres.UpdateDrinkRequest{
		Id:          req.ID,
		Title:       body.Title,
		Price:       body.Price,
		DrinkType:   postgres.DrinkType(body.DrinkType),
		Description: body.Description,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.UpdateDrink(ctx, request)

	if err != nil {
		log.Fatalf("could not update drink: %v", err)
		c.String(http.StatusInternalServerError, "Updating drink failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}

// / create gRPC delete drink request from HTTP request body
// / established, remove comment later
func (app *AppCore) DeleteDrink(c *gin.Context) {

	type Req struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req Req
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var request = &postgres.DeleteDrinkRequest{
		Id: req.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.DeleteDrink(ctx, request)

	if err != nil {
		log.Fatalf("could not delete drink: %v", err)
		c.String(http.StatusInternalServerError, "Deleting drink failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}

// / create gRPC get drink request from HTTP request body
// / established, remove comment later
func (app *AppCore) GetDrink(c *gin.Context) {

	type getBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req getBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.GetDrink(ctx, &postgres.GetDrinkRequest{
		Id: req.ID,
	})

	if err != nil {
		log.Fatalf("could not get drink: %v", err)
		c.String(http.StatusInternalServerError, "Getting bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}
