package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"log"
)

var DB *gorm.DB

func InitDBConnection() *gorm.DB {
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	DB, err := gorm.Open("mysql", dbConnectionString+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln(err)
	}

	DB.Exec("USE storefront")

	return DB
}

