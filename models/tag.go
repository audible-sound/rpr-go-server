package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // Creates a version 4 UUID
	Name     uuid.UUID `json:"name" gorm:"not null; unqiue"`
	PostTags []PostTag `json:"post_tag" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE, onDelete:CASCADE;"`
}
