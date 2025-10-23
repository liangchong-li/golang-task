package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	username, password, host := "root", "admin", "192.168.2.2"
	port := 3306
	Dbname := "gorm"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//ctx := context.Background()

	db.AutoMigrate(&Product{})
}
