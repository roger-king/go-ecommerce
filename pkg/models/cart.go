package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model

	User User `gorm:"foreignkey:UserID;association_foreignkey:Refer"`
	Products []Product `gorm:"many2many:carts_products;"`
}
