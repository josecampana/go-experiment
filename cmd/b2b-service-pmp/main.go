package main

import (
	"log"
	"net/http"

	logger "github.com/ingka-group-digital/b2b-service-pmp/pkg"

	env "github.com/ingka-group-digital/b2b-service-pmp/configs"
	routes "github.com/ingka-group-digital/b2b-service-pmp/internal/routes"
)

func main() {
	
	config := env.Get()
	router := routes.InitRoutes()

	log.Println("Listening requests\t➜ http://localhost:" + config.Port)
	logger.Init()
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
