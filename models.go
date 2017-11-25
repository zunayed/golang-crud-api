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

func createProduct(DB *gorm.DB, code string, price uint) (uint, error) {
	p := Product{Code: code, Price: price}

	err := DB.Create(&p).Error
	if err != nil {
		return 0, err
	}

	return p.ID, nil
}

func deleteProduct(DB *gorm.DB, price uint) error {
	var p Product
	DB.Where("id = ?", price).First(&p)
	err := DB.Delete(&p).Error
	if err != nil {
		return err
	}

	return nil
}
