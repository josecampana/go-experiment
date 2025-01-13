package controllers

import (
	"fmt"
	"net/http"
	env "github.com/ingka-group-digital/b2b-service-pmp/configs"
)

func Hej(w http.ResponseWriter, r *http.Request) {
	name := env.Get().AppName
	fmt.Fprintf(w, "Hej from %s (golang ver)", name)
}
