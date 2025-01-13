package controllers

import (
	"fmt"
	"net/http"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
)

func Hej(w http.ResponseWriter, r *http.Request) {
	name := config.Get().AppName
	fmt.Fprintf(w, "Hej from %s (golang ver)", name)
}
