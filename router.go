package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/go-ecommerce/heartbeat"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/api/healthCheck",
		heartbeat.HealthCheckController,
	},
	//Route {
	//	"GetProducts",
	//	"GET",
	//	"/api/store",
	//	Store.GetProductController,
	//},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
