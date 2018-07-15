package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Application Router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Query *Query
}

type Query struct {
	key string
	value string
}

type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

type Routes []Route

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/api/healthCheck",
		HealthCheckController,
		nil,
	},
	Route{
		"Login",
		"POST",
		"/api/login",
		AuthenticateController,
		nil,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/user",
		CreateUserController,
		nil,
	},
	Route{
		"CreateUser",
		"GET",
		"/api/user",
		FindUserByEmailController,
		&Query {
			"email",
			"{id:[0-9]+}",
		},
	},
	Route{
		"GetProducts",
		"GET",
		"/api/store",
		FindProductsController,
		nil,
	},
	Route{
		"CreateProducts",
		"POST",
		"/api/store",
		CreateProductsController,
		nil,
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

		if route.Query != nil {
			router.Queries(route.Query.key, route.Query.value)
		}
	}

	return router
}
