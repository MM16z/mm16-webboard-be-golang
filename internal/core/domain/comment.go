package domain

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint
	UserID    uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time

	Post User `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
