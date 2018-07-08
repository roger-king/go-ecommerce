package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"

	"github.com/roger-king/go-ecommerce/server"
	"github.com/roger-king/go-ecommerce/db"
	"github.com/roger-king/go-ecommerce/models"
)

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
	db := db.InitDBConnection()

	defer db.Close()

	db.AutoMigrate(&models.Product{})


	if port == "" {
		log.Fatalln("$PORT is not defined")
	}

	router := server.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Errorln(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
	log.Infoln("Application Started on http://localhost:8000")
}
