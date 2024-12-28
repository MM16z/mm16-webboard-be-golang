package postgres

import (
	"log"

	"github.com/MM16z/mm16-webboard-be-golang/internal/config"
	"github.com/MM16z/mm16-webboard-be-golang/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func InitDB() *DB {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	db, err := gorm.Open(postgres.Open(config.DBConfig.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.Exec(`SET search_path TO "mm16-webboard-golang"`).Error
	if err != nil {
		log.Fatalf("Error to set schema: %v", err)
	}

	err = db.Exec(`CREATE SCHEMA IF NOT EXISTS "mm16-webboard-golang"`).Error
	if err != nil {
		log.Fatalf("Error to create schema: %v", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.Comment{}, &domain.PostLiked{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return &DB{DB: db}
}
