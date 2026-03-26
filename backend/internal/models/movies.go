package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	Duration    string `json:"duration" db:"duration"`
	Language    string `json:"language" db:"language"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
