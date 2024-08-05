package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	Name      string    `gorm:"size:24;uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Blogs     []Blog    `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:(gen_random_uuid())"`
	Title     string    `gorm:"not null;index"`
	Text      string    `gorm:"not null"`
	AuthorID  uuid.UUID `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
