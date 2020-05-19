package model

import (
	"gbstore/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func SetProduct() {
	var err error
	db, err = gorm.Open(config.Conf.DB.Type, config.Conf.DB.DSN)
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{})
	db.LogMode(true)
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Cart{})
	db.AutoMigrate(&User{})
}
