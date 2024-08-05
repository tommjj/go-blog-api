package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

var jwtMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

type CustomClaims struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	jwt.RegisteredClaims
}

type JWTService struct {
	key      []byte
	keyFunc  func(token *jwt.Token) (interface{}, error)
	duration time.Duration
}

func NewJWTTokenService(conf config.Auth) (*JWTService, error) {
	duration, err := time.ParseDuration(conf.Duration)
	if err != nil {
		return nil, err
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrInvalidToken
		}

		return []byte(conf.SecretKey), nil
	}

	return &JWTService{
		key:      []byte(conf.SecretKey),
		keyFunc:  keyFunc,
		duration: duration,
	}, nil
}

func (j *JWTService) CreateToken(user *domain.User) (string, error) {
	claims := jwt.NewWithClaims(jwtMethod, CustomClaims{
		user.ID,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	})

	str, err := claims.SignedString(j.key)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return str, nil
}

func (j *JWTService) VerifyToken(tokenString string) (*domain.TokenPayload, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, j.keyFunc)

	switch {
	case token.Valid:
		return &domain.TokenPayload{
			ID:   claims.ID,
			Name: claims.Name,
		}, nil
	case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, domain.ErrInvalidToken
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, domain.ErrExpiredToken
	default:
		return nil, err
	}
}
