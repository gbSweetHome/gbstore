package model

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserId    uint
	ProductId uint
}

type UserCart struct {
	ID        uint
	Name      string
	Price     float64
	ImageUrl  string
	ProductId uint
}

func GetUserCart(userID uint) ([]UserCart, error) {
	var res []UserCart
	err := db.Table("cart").
		Select("cart.id, product.name, product.price,product.image_url,product.id as product_id").
		Joins("left join product on product.id = cart.product_id").
		Joins("left join user on user.id = cart.user_id").
		Where("cart.user_id = ?", userID).
		Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddUserCart(userID, productID uint) error {
	var cart = &Cart{
		UserId:    userID,
		ProductId: productID,
	}
	err := db.Create(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserCart(userID uint, cartID []uint) error {
	var cart = &Cart{}
	err := db.Unscoped().Where("id in (?)", cartID).Delete(cart).Error
	if err != nil {
		return err
	}
	return nil
}
