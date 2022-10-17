package server

import (
	"app/main/logger"
	"app/main/postgres"
	"app/main/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppCore struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`

	DataServiceHost string `json:"data_access_service_host"`
	DataServicePort uint16 `json:"data_access_service_port"`

	Logger *logger.LoggerCore
	Router *gin.Engine

	dataService postgres.BarMapServiceClient
}

func (app *AppCore) InitDataServiceGrpcConnection() {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", app.DataServiceHost, app.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	app.dataService = postgres.NewBarMapServiceClient(conn)
}

func Start(configPath string, logger *logger.LoggerCore) {

	var app AppCore
	utils.ParseJsonConfig(configPath, &app)
	app.Logger = logger

	app.InitDataServiceGrpcConnection()
	app.InitRouter()

	app.Router.Run(fmt.Sprintf("%s:%d", app.Host, app.Port))
}
