package service

import (
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type AuthService struct {
	tk *ports.ITokenService
}

func NewAuthService(token *ports.ITokenService) *AuthService {
	return &AuthService{
		tk: token,
	}
}

// func (as *AuthService) Login(ctx context.Context, username, password string) (string, error) {

// }
