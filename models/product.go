package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model

	Name string `json:"name"`
}

func AllProducts() ([]Product, error) {
	var products []Product

	defer db.Close()

	err := db.Find(&products).Error

	return products, err
}

func CreateProduct(product Product) (*Product, error) {
	var err error
	err = db.Create(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, err
}
