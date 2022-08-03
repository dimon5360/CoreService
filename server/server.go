package server

import (
	"app/main/kafka"
	"app/main/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ingredient struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Count int32  `json:"count"`
}

type Drink struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Price       int32  `json:"price"`
	Description string `json:"description"`
	Type        int32  `json:"type"`
	// Ingredient  []Ingredient `json:"ingredient"`
}

// **************** API:
// GET /drinks
// GET /drinks/<id>
// POST /drinks
// DELETE /drinks/<id>
// UPDATE /drinks/<id>

var drinks = []Drink{
	{Id: 1, Title: "Drink1", Price: 5699, Description: "Alcohol drink 1", Type: 1},
	{Id: 2, Title: "Drink2", Price: 1799, Description: "Alcohol drink 1", Type: 1},
	{Id: 3, Title: "Drink3", Price: 3999, Description: "Alcohol drink 1", Type: 1},
}

func GetDrinks(ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, drinks)
}

type ServerConfig struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`

	kafkaHandler *kafka.Handler
}

func SetupRouting(router *gin.Engine) {
	router.GET("/drinks", GetDrinks)
	// todo
}

func Start(confFileName string) {

	var config ServerConfig
	utils.ParseJsonConfig(confFileName, &config)

	router := gin.Default()
	SetupRouting(router)
	go func() { router.Run(fmt.Sprintf("%s:%d", config.Host, config.Port)) }()

	config.kafkaHandler = kafka.InitKafkaHandler("config/kafka.json")

	config.kafkaHandler.StartKafkaHandler()

}
