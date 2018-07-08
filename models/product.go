package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	Name string `json:"name"`
}

func (db *DB) AllProducts() ([]Product, error) {
	var products []Product

	defer db.Close()

	err := db.Find(&products).Error

	return products, err
}

func (db *DB) CreateProduct(product Product) (*Product, error) {
	err := db.Create(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, err
}
