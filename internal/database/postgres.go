package database

import (
	"fmt"
	"log"
	"movie-recommendation/internal/config"
	"movie-recommendation/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.Genre{},
		&models.Movie{},
		&models.User{},
		&models.Rating{},
	)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	log.Println("Database migration completed")
	return db, nil
}
