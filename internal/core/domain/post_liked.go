package domain

import (
	"time"

	"gorm.io/gorm"
)

type PostLiked struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Post Post `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (PostLiked) TableName() string {
	return "post_liked"
}

func (PostLiked) BeforeCreate(tx *gorm.DB) error {
	err := tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS post_liked_un ON post_liked (user_id, post_id)").Error
	if err != nil {
		return err
	}
	return nil
}
