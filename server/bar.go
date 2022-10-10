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

// GET request - form to create POST request for new bar
func (app *AppConfig) CreateBarForm(c *gin.Context) {

	c.HTML(http.StatusOK, "create_bar.html", gin.H{
		"content": "This is an index page...",
	})
}

// / create gRPC create bar request from HTTP request body
// / established, remove comment later
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

	type BarResponse struct {
		Id          string
		Title       string
		Address     string
		Description string
	}

	resp := make([]BarResponse, 0)
	resp = append(resp, BarResponse{
		Id:          r.Id,
		Title:       r.Title,
		Address:     r.Address,
		Description: r.Description,
	})

	c.JSON(http.StatusCreated, resp)
}

// / create gRPC update bar request from HTTP request body
// / established, remove comment later
func (app *AppConfig) UpdateBar(c *gin.Context) {

	type updateBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req updateBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	body := BarBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.UpdateBar(ctx, &postgres.UpdateBarRequest{
		Id:          req.ID,
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

// / create gRPC delete bar request from HTTP request body
// / established, remove comment later
func (app *AppConfig) DeleteBar(c *gin.Context) {
	type deleteBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req deleteBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.DeleteBar(ctx, &postgres.DeleteBarRequest{
		Id: req.ID,
	})

	if err != nil {
		log.Fatalf("could not create bar: %v", err)
		c.String(http.StatusInternalServerError, "Creating bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusCreated, r.String())
}

// / create gRPC get bar request from HTTP request body
// / established, remove comment later
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
