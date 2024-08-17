package domain

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text,omitempty"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
