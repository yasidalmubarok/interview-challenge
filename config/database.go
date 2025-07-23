package config

import (
	"challenge-interview/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	LoadEnv()
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		panic("DB_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	// Set the connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance: " + err.Error())
	}
	
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)
	if err := db.AutoMigrate(&entity.Car{}, &entity.Order{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	
	return db
}
