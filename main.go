package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"

	"github.com/roger-king/go-ecommerce/models"
	"github.com/roger-king/go-ecommerce/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	db gorm.DB
}

// Application Router
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
		controllers.HealthCheckController,
	},
	Route {
		"GetProducts",
		"GET",
		"/api/store",
		controllers.FindProductsController,
	},
	Route {
		"CreateProducts",
		"POST",
		"/api/store",
		controllers.CreateProductsController,
	},
}

func (app *App) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Infoln(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}


func main() {
	port := os.Getenv("PORT")

	// DB Connection
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, dbError := models.InitDB(dbConnectionString)

	if dbError != nil {
		log.Panicf("Error connecting to database %s", dbError)
	}

	defer db.Close()

	db.AutoMigrate(&models.Product{})

	app := &App{db: db}

	if port == "" {
		log.Panic("$PORT is not defined")
	}

	router := NewRouter(app)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Errorln(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
	log.Infoln("Application Started on http://localhost:8000")
}
