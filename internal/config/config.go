package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig *DBConfig
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type DBConfig struct {
	DSN string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, err
	}

	db := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dsn := &DBConfig{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=mm16-webboard-golang",
			db.Host,
			db.User,
			db.Password,
			db.DBName,
			db.Port,
		),
	}

	return &Config{
		dsn,
	}, nil
}
