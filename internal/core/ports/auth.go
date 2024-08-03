package ports

import (
	"context"

	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type TokenService interface {
	CreateToken(user *domain.User) (string, error)

	VerifyToken(token string) (*domain.TokenPayload, error)
}
