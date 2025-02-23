package models

import (
	"github.com/google/uuid"
)

type PostTag struct {
	PostID uuid.UUID `json:"post_id" gorm:"primaryKey;"`
	TagID  uuid.UUID `json:"tag_id" gorm:"primaryKey;"`
	Post   Post      `json:"post" gorm:"foreignKey:PostID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
	Tag    Tag       `json:"tag" gorm:"foreignKey:TagID; constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
}
