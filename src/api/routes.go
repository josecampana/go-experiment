package routes

import (
	"b2b-service-pmp/env"
	controllers "b2b-service-pmp/src/api/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logRoute(basePath string) func(method string, path string) {
	return func(method string, path string) {
		log.Println(fmt.Sprintf("Added route \t➜ %s\t%s%s", method, basePath, path))
	}
}

func InitRoutes() *mux.Router {
	config := env.Get()
	router := mux.NewRouter()
	apiRouter := router.PathPrefix(config.BasePath).Subrouter()

	// solves the problem of trailing slash
	router.HandleFunc(config.BasePath, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.BasePath+"/", http.StatusMovedPermanently)
	})

	//redirection to basepath
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		protocol := "http"
		if r.Header.Get("X-Forwarded-Proto") == "https" {
			protocol = "https"
		}

		host := r.Host
		http.Redirect(w, r, fmt.Sprintf("%s://%s/%s", protocol, host, config.BasePath), http.StatusMovedPermanently)
	}).Methods("GET")

	//api routes
	apiRouter.HandleFunc("/", controllers.Hej).Methods("GET")
	apiRouter.HandleFunc("/stores", controllers.GetStores).Methods("GET")
	apiRouter.HandleFunc("/stores", controllers.CreateStore).Methods("POST")
	apiRouter.HandleFunc("/stores/{storeId}", controllers.GetStoreById).Methods("GET")
	apiRouter.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")

	logRoute(config.BasePath)("GET", "/")
	logRoute(config.BasePath)("GET", "/stores")
	logRoute(config.BasePath)("POST", "/stores")
	logRoute(config.BasePath)("GET", "/stores/{storeId}")
	logRoute(config.BasePath)("GET", "/product/{id}")

	return router
}
