package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/rogr-king/go-ecommerce/app"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error

	engine, err = xorm.NewEngine("mysql", "localhost/devdb?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT is not defined")
	}

	router := app.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}