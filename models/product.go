package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	name string `json:"name"`
}
