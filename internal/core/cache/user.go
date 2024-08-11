package cache

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type UserCacheService struct {
	cache    ports.ICacheRepository // Cache ICacheRepository
	duration time.Duration          // cache storage time
}

func NewUserCacheService(cache ports.ICacheRepository, duration time.Duration) *UserCacheService {
	return &UserCacheService{
		cache,
		duration,
	}
}

func (ucs *UserCacheService) SetUser(ctx context.Context, user *domain.User) error {
	bytes, err := marshal(user)
	if err != nil {
		return err
	}

	return ucs.cache.Set(ctx, user.ID.String(), bytes, ucs.duration)
}

func (ucs *UserCacheService) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	bytes, err := ucs.cache.Get(ctx, id.String())
	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	err = unmarshal(bytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ucs *UserCacheService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ucs.cache.Delete(ctx, id.String())
}
