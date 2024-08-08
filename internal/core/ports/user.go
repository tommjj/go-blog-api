package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type IUserRepository interface {
	// GetUserByID select a user by id
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	// GetUserByName select a user by name
	GetUserByName(ctx context.Context, name string) (*domain.User, error)
	// CreateUser insert a new user into the database
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// UpdateUser update a user, only update non-zero fields by default
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// UpdateUserByMap update a user, update by map data
	UpdateUserByMap(ctx context.Context, id uuid.UUID, data *map[string]interface{}) (*domain.User, error)
	// DeleteUser delete a user
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
