package services

import (
	"github.com/roger-king/go-ecommerce/db"
	"github.com/roger-king/go-ecommerce/models"
)

var dbConnection = db.InitDBConnection()


func GetProducts() ([]models.Product, error) {
	var products []models.Product

	defer dbConnection.Close()

	err := dbConnection.Find(&products).Error

	return products, err
}

func CreateProduct(product models.Product) bool {
	dbConnection.Create(&product)

	return true
}
