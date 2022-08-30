package server

import (
	"app/main/postgres"
	"app/main/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppConfig struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`

	DataServiceHost string `json:"data_access_service_host"`
	DataServicePort uint16 `json:"data_access_service_port"`

	dataService postgres.BarMapServiceClient
}

func InitDataServiceGrpcConnection(app *AppConfig) postgres.BarMapServiceClient {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", app.DataServiceHost, app.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return postgres.NewBarMapServiceClient(conn)
}

func Start(confFileName string) {

	var app AppConfig
	utils.ParseJsonConfig(confFileName, &app)

	router := gin.Default()

	app.dataService = InitDataServiceGrpcConnection(&app)

	SetupRouting(router, &app)

	router.Run(fmt.Sprintf("%s:%d", app.Host, app.Port))
}
