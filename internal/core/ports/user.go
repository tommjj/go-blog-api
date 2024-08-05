package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUserByName(ctx context.Context, name string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUserByMap(ctx context.Context, id uuid.UUID, data *map[string]interface{}) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
