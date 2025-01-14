package routes

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	controllers "github.com/ingka-group-digital/b2b-service-pmp/internal/controllers"
	middlewares "github.com/ingka-group-digital/b2b-service-pmp/internal/middlewares"

	"github.com/gorilla/mux"
)

func logRoute(basePath string) func(method string, path string) {
	return func(method string, path string) {
		log.Println(fmt.Sprintf("Added route \t➜ %s\t%s%s", method, basePath, path))
	}
}

func InitRoutes() *mux.Router {
	config := config.Get()
	router := mux.NewRouter()
	router.Use(middlewares.ContextMW)

	apiRouter := router.PathPrefix(config.BasePath).Subrouter()

	// solves the problem of trailing slash
	router.HandleFunc(config.BasePath, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.BasePath+"/", http.StatusTemporaryRedirect)
	})

	//redirection to basepath
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		protocol := "http"
		if r.Header.Get("X-Forwarded-Proto") == "https" {
			protocol = "https"
		}

		host := r.Host
		http.Redirect(w, r, fmt.Sprintf("%s://%s/%s", protocol, host, config.BasePath), http.StatusTemporaryRedirect)
	}).Methods("GET")

	//api routes
	apiRouter.HandleFunc("/", controllers.Hej).Methods("GET")
	// apiRouter.HandleFunc("/stores", controllers.GetStores).Methods("GET")
	// apiRouter.HandleFunc("/stores", controllers.CreateStore).Methods("POST")
	// apiRouter.HandleFunc("/stores/{storeId}", controllers.GetStoreById).Methods("GET")

	//range experiment
	apiRouter.HandleFunc("/v3/{retailUnit}/{language}/products", controllers.GetProducts).Methods("GET")

	logRoute("")("GET", "/")
	logRoute(config.BasePath)("GET", "/")
	logRoute(config.BasePath)("GET", "/v3/{retailUnit}/{language}/products")
	// logRoute(config.BasePath)("GET", "/stores")
	// logRoute(config.BasePath)("POST", "/stores")
	// logRoute(config.BasePath)("GET", "/stores/{storeId}")
	// logRoute(config.BasePath)("GET", "/product/{id}")

	return router
}
