package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"` // Creates a version 4 UUID
	PostPictureURL  string     `json:"post_picture_url" gorm:"not null"`
	PostTitle       string     `json:"post_title" gorm:"not null"`
	PostDescription string     `json:"post_description" gorm:"not null"`
	Location        string     `json:"location" gorm:"not null"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoCreateTime"`
	UserID          uuid.UUID  `json:"user_id" gorm:"not null; index"`
	PostLikes       []PostLike `json:"post_likes" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	Comments        []Comment  `json:"comments" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	PostTags        []PostTag  `json:"post_tag" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
}

func init() {
	Register(&Post{})
}
