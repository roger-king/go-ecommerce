package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model

	User     User      `gorm:"foreignkey:UserID;association_foreignkey:Refer"`
	Products []Product `gorm:"many2many:carts_products;"`
}

func createCart(cart *Cart) (*Cart, error) {
	var err error
	err = db.Create(&cart).Error

	if err != nil {
		return nil, err
	}

	createdCart := Cart{
		User:     cart.User,
		Products: cart.Products,
	}

	return &createdCart, nil
}
