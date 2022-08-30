package server

import (
	"app/main/postgres"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DrinkBody struct {
	Title       string
	Price       string
	DrinkType   int32
	Description string
	BarId       int32
	Ingredients []IngredientBody
}

func (app *AppConfig) CreateDrink(c *gin.Context) {

	body := DrinkBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request = &postgres.CreateDrinkRequest{
		Title:       body.Title,
		Price:       body.Price,
		DrinkType:   postgres.DrinkType(body.DrinkType),
		Description: body.Description,
		BarId:       fmt.Sprintf("%d", body.BarId),
		Ingredients: []*postgres.CreateIngredientRequest{},
	}

	for _, ingredient := range body.Ingredients {
		var tmp = &postgres.CreateIngredientRequest{
			Title:   ingredient.Title,
			Amount:  fmt.Sprintf("%d", ingredient.Amount),
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

	c.String(http.StatusCreated, r.String())
}

func (app *AppConfig) GetDrink(c *gin.Context) {

	type getBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req getBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Printf("Got request for drink with ID = %s\n", req.ID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.GetDrink(ctx, &postgres.GetDrinkRequest{
		Id: req.ID,
	})

	if err != nil {
		log.Fatalf("could not get drink: %v", err)
		c.String(http.StatusInternalServerError, "getting bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}
