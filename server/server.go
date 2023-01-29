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

	StorageServiceHost string `json:"storage_service_host"`

	Logger *logger.LoggerCore
	Router *gin.Engine

	dataService postgres.BarMapServiceClient
}

func (app *AppCore) InitStorageServiceGrpcConnection() {

	conn, err := grpc.Dial(app.StorageServiceHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	app.dataService = postgres.NewBarMapServiceClient(conn)
}

func Start(configPath string, loggerConfigPath string) {

	var app AppCore
	utils.ParseJsonConfig(configPath, &app)
	app.Logger = logger.Init(loggerConfigPath)

	app.InitStorageServiceGrpcConnection()
	app.InitRouter()

	app.Logger.Write(fmt.Sprintf("Start listening %s", app.Host))

	app.Router.Run(app.Host)
}
