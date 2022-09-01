package server

import (
	"app/main/postgres"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BarBody struct {
	Title       string
	Address     string
	Description string
}

func (app *AppConfig) CreateBar(c *gin.Context) {

	body := BarBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.CreateBar(ctx, &postgres.CreateBarRequest{
		Title:       body.Title,
		Address:     body.Address,
		Description: body.Description,
	})

	if err != nil {
		log.Fatalf("could not create bar: %v", err)
		c.String(http.StatusInternalServerError, "Creating bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusCreated, r.String())
}

func (app *AppConfig) GetBar(c *gin.Context) {

	type getBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req getBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Printf("Got request for bar with ID = %s\n", req.ID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.GetBar(ctx, &postgres.GetBarRequest{
		Id: req.ID,
	})

	if err != nil {
		log.Fatalf("could not get bar: %v", err)
		c.String(http.StatusInternalServerError, "getting bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}
