package service

import (
	"context"

	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
	"github.com/tommjj/go-blog-api/internal/core/util"
)

type AuthService struct {
	tk   ports.ITokenService
	repo ports.IUserRepository
}

func NewAuthService(token ports.ITokenService, userRepo ports.IUserRepository) ports.IAuthService {
	return &AuthService{
		tk:   token,
		repo: userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := as.repo.GetUserByName(ctx, username)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := as.tk.CreateToken(user)
	if err != nil {
		return "", domain.ErrInternal
	}

	return token, nil
}
