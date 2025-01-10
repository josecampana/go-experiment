package main

import (
	"b2b-service-pmp/env"
	routes "b2b-service-pmp/src/api"
	"log"
	"net/http"
)

func main() {
	config := env.Get()
	router := routes.InitRoutes()

	log.Println("Listening requests\t➜ http://localhost:" + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
