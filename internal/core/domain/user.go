package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Email          string `gorm:"unique"`
	Password       string
	Username       string `gorm:"unique"`
	RefreshToken   sql.NullString
	ProfilePicture sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
