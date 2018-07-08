package models

import "github.com/jinzhu/gorm"

type DataStore interface {
	AllProducts() ([]*Product, error)
	CreateProduct(product Product)
}

type DB struct {
	*gorm.DB
}

func InitDB(dsn string) (*DB, error) {
	db, err := gorm.Open("mysql", dsn+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
