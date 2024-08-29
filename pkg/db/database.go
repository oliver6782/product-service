package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"product-service/internal/config"
	"product-service/internal/model" 
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Add this line to auto-migrate your models
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
