package main

import (
	"app/main/server"
	"log"
)

// build v.0.0.4 from 16.09.2022
const (
	BUILD = 4
	MINOR = 0
	MAJOR = 0
)

func main() {

	log.Printf("Start Web server service v.%d.%d.%d.", MAJOR, MINOR, BUILD)

	server.Start("config/server.json")
}
