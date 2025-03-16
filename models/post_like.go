package models

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	UserID  uuid.UUID `json:"user_id" gorm:"primaryKey;"`
	PostID  uuid.UUID `json:"post_id" gorm:"primaryKey;"`
	User    User      `json:"user" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	Post    Post      `json:"post" gorm:"foreignKey:PostID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	LikedAt time.Time `json:"liked_at" gorm:"autoCreateTime"`
}

func init() {
	Register(&PostLike{})
}
