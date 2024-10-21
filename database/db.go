package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB(config *Config) *gorm.DB {
	var err error
	var db *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	return db
}

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	DB = ConnectDB(&config)
}
