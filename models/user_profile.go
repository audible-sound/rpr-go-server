package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"` // Creates a version 4 UUID
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;not null;unique"`
	User            User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	BirthDate       time.Time `json:"birth_date" gorm:"not null; check:birth_date >= '1900-01-01' AND birth_date <= now()"` // Birthdates must be valid
	Bio             string    `json:"bio" gorm:"type:varchar(500);check:length(bio) >= 10;not null"`                        // Bio length must be between 10 to 500
	ProfileImageURL string    `json:"profile_image_url" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func init() {
	Register(&UserProfile{})
}
