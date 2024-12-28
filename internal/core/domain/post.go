package domain

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	UserName  string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
	PostLiked []PostLiked
	User      User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
