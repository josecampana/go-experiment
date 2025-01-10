package controllers

import (
	"fmt"
	"net/http"
)

func Hej(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hej from b2b-service-pmp")
}
