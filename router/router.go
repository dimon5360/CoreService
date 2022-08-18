package router

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

func SetupRouting(router *gin.Engine, app *appConfig) {

	router.POST("/createbar", app.CreateBar)
}

// curl -i -X POST -H 'Content-Type: application/json'
// -d '{"id": "1234", "title": "Goodson", "address": "Zelenograd",
// "description": "bar"}' http://localhost:40400/createbar

func (app *appConfig) CreateBar(c *gin.Context) {

	log.Println("Create bar called")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Query("id")
	title := c.DefaultQuery("title", "Unknown")
	address := c.PostForm("address")
	description := c.PostForm("description")

	r, err := app.dataService.CreateBar(ctx, &postgres.CreateBarRequest{
		Id:          id,
		Title:       title,
		Address:     address,
		Description: description,
		Drinks:      []*postgres.CreateDrinkRequest{},
	})

	if err != nil {
		log.Fatalf("could not create bar: %v", err)
		c.String(http.StatusInternalServerError, "Creating bar failed")
	}
	c.String(http.StatusCreated, r.String())
}

func InitDataServiceGrpcConnection(app *appConfig) postgres.BarMapServiceClient {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", app.DataServiceHost, app.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return postgres.NewBarMapServiceClient(conn)
}

	var config Config

	var app appConfig
	utils.ParseJsonConfig(confFileName, &app)

	router := gin.Default()

	app.dataService = InitDataServiceGrpcConnection(&app)

	SetupRouting(router, &app)

	router.Run(fmt.Sprintf("%s:%d", app.Host, app.Port))
}
