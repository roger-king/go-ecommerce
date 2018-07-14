package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type DataStore interface {
	AllProducts() ([]*Product, error)
	CreateProduct(product Product)
}

var db *gorm.DB

func InitDB(dsn string) {
	var err error
	db, err = gorm.Open("mysql", dsn+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Panicln(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Panicf("Error connecting to database: %s", err)
	}

	// TODO: Before Production we want to dynamically set all of this.
	db.LogMode(true)

	db.AutoMigrate(&Product{}, &User{})
}
