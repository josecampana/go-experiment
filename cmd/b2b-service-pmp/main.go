package main

import (
	"log"
	"net/http"

	utils "github.com/ingka-group-digital/b2b-service-pmp/internal/utils"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	routes "github.com/ingka-group-digital/b2b-service-pmp/internal/routes"
)

func main() {
	config := config.Get()
	router := routes.InitRoutes()

	log.Println("Listening requests\t➜ http://localhost:" + config.Port)

	utils.InitLogger(config.AppName)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))

}
