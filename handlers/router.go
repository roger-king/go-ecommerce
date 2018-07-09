package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/go-ecommerce/envs"
)

// Application Router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Handler struct {
	*envs.Env
	H func(e *envs.Env, w http.ResponseWriter, r *http.Request) error
}

type Routes []Route

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/api/healthCheck",
		HealthCheckController,
	},
	Route{
		"GetProducts",
		"GET",
		"/api/store",
		FindProductsController,
	},
	Route{
		"CreateProducts",
		"POST",
		"/api/store",
		CreateProductsController,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
