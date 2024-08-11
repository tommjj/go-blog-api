package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type UserService struct {
	ur    ports.IUserRepository
	cache ports.ICacheRepository
}

func NewUserService(userRepo ports.IUserRepository, cache ports.ICacheRepository) *UserService {
	return &UserService{
		ur:    userRepo,
		cache: cache,
	}
}

func (us *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := us.ur.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
