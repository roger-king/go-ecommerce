package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataStore interface {
	AllProducts() ([]*Product, error)
	CreateProduct(product Product)
}

type DB struct {
	*gorm.DB
}

type Connection struct {
	*DB
	dsn string
}

func InitDB(conn *Connection) (*DB, error) {
	db, err := gorm.Open("mysql", conn.dsn+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
