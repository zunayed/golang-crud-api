package main

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func getProduct(DB *gorm.DB, id int) (Product, error) {
	var product Product
	err := DB.First(&product, id).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
