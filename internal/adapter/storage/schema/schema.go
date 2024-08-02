package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	Name      string    `gorm:"size:24"`
	Password  string
	Blogs     []Blog `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	Title     string
	Text      string
	AuthorID  uuid.UUID
	Author    User `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
