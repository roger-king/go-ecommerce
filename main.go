package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/rogr-king/go-ecommerce/app"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()


	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT is not defined")
	}

	router := app.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}