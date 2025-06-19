package domain

import "github.com/google/uuid"

type Movie struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Poster      string    `gorm:"not null" json:"poster"`
	Trailer     string    `gorm:"not null" json:"trailer"`
	Actors      []string  `gorm:"type:text[];not null" json:"actors"`
	Genres      []string  `gorm:"type:text[];not null" json:"genres"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}
