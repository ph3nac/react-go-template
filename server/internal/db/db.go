package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
