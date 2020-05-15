package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Name        string
	ImageUrl    string
	Price       float64
	MonthlySale string `gorm:"default:0"`
	Describe    string `gorm:"type:text"`
	Rate        *int64 `gorm:"default:0"`
}

func AddProduct(name string, describe string, price float64, imageUrl string) (uint, error) {
	productItem := &Product{Name: name, Describe: describe, Price: price, ImageUrl: imageUrl}
	err := db.Create(productItem).Error
	if err != nil {
		return 0, err
	}
	return productItem.ID, nil
}

func GetProducts(pageNum int, pageSize int) ([]Product, int, error) {
	var products []Product
	var total int = 0
	err := db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Model(&Product{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func GetProduct(id uint) (*Product, error) {
	product := &Product{}
	product.ID = id
	err := db.Where("id=?", id).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func DeleteProduct(id uint) error {
	product := &Product{}
	product.ID = id
	err := db.Model(&Product{}).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductsByIDS(ids []uint) ([]*Product, error) {
	var products []*Product
	err := db.Where("id in (?)", ids).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
