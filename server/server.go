package server

import (
	"app/main/postgres"
	"app/main/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type appConfig struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`

	DataServiceHost string `json:"data_access_service_host"`
	DataServicePort uint16 `json:"data_access_service_port"`

	dataService postgres.BarMapServiceClient
}

type Body struct {
	Title       string
	Address     string
	Description string
}

func SetupRouting(router *gin.Engine, app *appConfig) {

	// POST requests
	router.POST("/createbar", app.CreateBar)

	// GET requests
	router.GET("/bar/:id", app.GetBar)
}

func (app *appConfig) CreateBar(c *gin.Context) {

	type Body struct {
		Title       string
		Address     string
		Description string
	}

	body := Body{}
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
		Drinks:      []*postgres.CreateDrinkRequest{},
	})

	if err != nil {
		log.Fatalf("could not create bar: %v", err)
		c.String(http.StatusInternalServerError, "Creating bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusCreated, r.String())
}

func (app *appConfig) GetBar(c *gin.Context) {
	type getBarRequest struct {
		ID string `uri:"id" binding:"required,min=1"`
	}

	var req getBarRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// log.Printf("Got request for bar with ID = %s\n", req.ID)
	// c.String(http.StatusOK, req.ID)

	// if err := c.BindJSON(&req); err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	log.Printf("Got request for bar with ID = %s\n", req.ID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := app.dataService.GetBar(ctx, &postgres.GetBarRequest{
		Id: req.ID,
	})

	if err != nil {
		log.Fatalf("could not create bar: %v", err)
		c.String(http.StatusInternalServerError, "Creating bar failed")
	}

	// TODO: serialize to JSON
	c.String(http.StatusOK, r.String())
}

func InitDataServiceGrpcConnection(app *appConfig) postgres.BarMapServiceClient {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", app.DataServiceHost, app.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return postgres.NewBarMapServiceClient(conn)
}

func Start(confFileName string) {

	var app appConfig
	utils.ParseJsonConfig(confFileName, &app)

	router := gin.Default()

	app.dataService = InitDataServiceGrpcConnection(&app)

	SetupRouting(router, &app)

	router.Run(fmt.Sprintf("%s:%d", app.Host, app.Port))
}
