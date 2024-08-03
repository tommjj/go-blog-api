package domain

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
