package main

import (
	"b2b-service-pmp/env"
	routes "b2b-service-pmp/src/api"
	logger "b2b-service-pmp/src/modules"
	"log"
	"net/http"
)

func main() {
	logger.Init()
	config := env.Get()
	router := routes.InitRoutes()

	log.Println("Listening requests\t➜ http://localhost:" + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
