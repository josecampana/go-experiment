package env

import (
	"os"
)

type Config struct {
	AppName	string
	Port     string //si empieza en mayúsculas, es público
	BasePath string
}

func Get() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return Config{Port: port, BasePath: "/range", AppName: "b2b-service-range"}
}
