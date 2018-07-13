package main

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	h "github.com/roger-king/go-ecommerce/handlers"
	"github.com/roger-king/go-ecommerce/models"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	port := os.Getenv("PORT")

	// DB Connection
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	models.InitDB(dbConnectionString)

	if port == "" {
		log.Panic("$PORT is not defined")
	}

	router := h.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Errorln(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
	log.Infoln("Application is running on port %s", port)
}
