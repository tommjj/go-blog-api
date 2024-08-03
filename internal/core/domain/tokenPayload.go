package domain

import "github.com/google/uuid"

type TokenPayload struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
