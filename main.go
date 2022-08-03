package main

import (
	"app/main/server"
	"log"
)

// build from 01.08.2022
const (
	BUILD = 1
	MINOR = 0
	MAJOR = 0
)

func main() {

	log.Printf("Start Web server service v.%d.%d.%d.", MAJOR, MINOR, BUILD)

	server.Start("config/server.json")
}
