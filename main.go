package main

import (
	"app/main/logger"
	"app/main/server"
	"log"
)

// build v.0.0.5 from 17.10.2022
const (
	BUILD = 5
	MINOR = 0
	MAJOR = 0
)

func main() {

	log.Printf("Start Web server service v.%d.%d.%d.", MAJOR, MINOR, BUILD)

	server.Start("config/server.json", logger.Init("config/kafka.json"))
}
