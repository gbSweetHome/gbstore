package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserId   uint
	Name     string
	Price    float64
	State    string
	ImageUrl string
}

func AddOrder(name string, price float64, imageUrl string, userId uint, status string) (uint, error) {
	orderItem := &Order{Name: name, Price: price, ImageUrl: imageUrl, UserId: userId, State: status}
	err := db.Create(orderItem).Error
	if err != nil {
		return 0, err
	}
	return orderItem.ID, nil
}

func GetOrders(pageNum int, pageSize int, userId uint) ([]Order, int, error) {
	var orders []Order
	var total int = 0
	err := db.Where("user_id=?", userId).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Model(&Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}
