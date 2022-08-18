package main

import (
	"app/main/kafka"
	"app/main/router"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// build from 01.08.2022
const (
	BUILD = 1
	MINOR = 0
	MAJOR = 0
)

type App struct {
	config       *router.Config
	kafkaHandler *kafka.Handler
	router       *gin.Engine
}

func (app *App) Start() {

	go func() {
		log.Fatal(app.router.Run(fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)))
	}()

	app.kafkaHandler.Start(app.router)
}

func main() {

	log.Printf("Start Web server service v.%d.%d.%d.", MAJOR, MINOR, BUILD)

	var app App

	app.kafkaHandler = kafka.Init("config/kafka.json")
	app.router, app.config = router.Init("config/server.json", app.kafkaHandler)

	app.Start()
}
