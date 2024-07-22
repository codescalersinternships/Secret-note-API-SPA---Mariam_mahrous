package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UniqueUrl      uuid.UUID `json:"unique_url"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	ExpirationDate string    `json:"expiration_date"`
	CurrentViews   int       `json:"current_views"`
	MaxViews       int       `json:"max_views"`
	UserID         uint      `json:"user_id"`
}

type CreateNoteInput struct {
	Title          string `json:"title"`
	Content        string `json:"content"`
	ExpirationDate string `json:"expiration_date"`
	MaxViews       int    `json:"max_views"`
}
