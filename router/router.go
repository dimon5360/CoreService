package router

import (
	"app/main/kafka"
	"app/main/utils"
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

type Router struct {
	kafka *kafka.Handler
}

var drinks = []Drink{
	{Id: 1, Title: "Drink1", Price: 5699, Description: "Alcohol drink 1", Type: 1},
	{Id: 2, Title: "Drink2", Price: 1799, Description: "Alcohol drink 1", Type: 1},
	{Id: 3, Title: "Drink3", Price: 3999, Description: "Alcohol drink 1", Type: 1},
}

func GetDrinks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, drinks)
}

// **************** API:
// GET /drinks
// GET /drinks/<id>
// POST /drinks
// DELETE /drinks/<id>
// UPDATE /drinks/<id>

func SetupRouting(router *gin.Engine, kafka *kafka.Handler) {
	router.GET("/drinks", GetDrinks)
	// todo
}

type Config struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

func Init(confFileName string, kafka *kafka.Handler) (*gin.Engine, *Config) {

	var config Config

	utils.ParseJsonConfig(confFileName, &config)

	router := gin.Default()
	SetupRouting(router, kafka)

	return router, &config
}
