package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type UserService struct {
	ur *ports.IUserRepository
}

func NewUserService(userRepo *ports.IUserRepository) *UserService {
	return &UserService{
		ur: userRepo,
	}
}

func (us *UserService) GetUserById(ctx context.Context, id uuid.UUID) {

}
