package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"` // Creates a version 4 UUID
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
	UserID    uuid.UUID `json:"user_id" gorm:"primaryKey;"`
	PostID    uuid.UUID `json:"post_id" gorm:"primaryKey;"`
	User      User      `json:"user" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	Post      Post      `json:"post" gorm:"foreignKey:PostID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
}

func init() {
	Register(&Comment{})
}
