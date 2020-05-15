package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string
	Avatar string
}

//暂时用name当作唯一id

func AddUser(name, avatar string) error {
	var user = &User{
		Name:   name,
		Avatar: avatar,
	}
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(name string) (*User, error) {
	var user User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Exist(name string) (bool, error) {
	var user []*User
	err := db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		return false, err
	}
	if len(user) == 0 {
		return false, nil
	}
	return true, nil
}
