package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/roger-king/go-ecommerce/models"
	"github.com/roger-king/go-ecommerce/server"
	log "github.com/sirupsen/logrus"
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

var DB *gorm.DB
var dbError error

func main() {
	port := os.Getenv("PORT")

	// DB Connection
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	DB, dbError = gorm.Open("mysql", dbConnectionString+"?charset=utf8&parseTime=True&loc=Local")

	if dbError != nil {
		log.Fatalln(dbError)
	}

	defer DB.Close()

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})

	if port == "" {
		log.Fatalln("$PORT is not defined")
	}

	router := server.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Errorln(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
	log.Infoln("Application Started on http://localhost:8000")
}
