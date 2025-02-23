package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // Creates a version 4 UUID
	Hash      string         `json:"-" gorm:"not null"`                                         // Hide from JSON Response
	Email     string         `json:"email" gorm:"unique;not null"`
	Username  string         `json:"username" gorm:"unique; not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // Enable soft delete, user record will be excluded when fetching data
	Posts     []Post         `json:"posts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	PostLikes []PostLike     `json:"post_likes" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	Comments  []Comment      `json:"comments" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
}
