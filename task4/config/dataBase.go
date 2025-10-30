package config

import (
	"fmt"
	"log"
	"os"
	"task4/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDataBase() {

	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "1234")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "golang_blog")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 迁移表结构
	err = DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		log.Fatal("AutoMigrate fail in initDataBase", err)
	}

	log.Println("initDataBase finished")
}

func getEnv(k string, defaultValue string) string {
	if value, exists := os.LookupEnv(k); exists {
		return value
	}
	return defaultValue
}
