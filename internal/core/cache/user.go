package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

var (
	userPrefix = "user"
)

// implement ports.IUserCacheService
type UserCacheService struct {
	cache    ports.ICacheRepository // Cache ICacheRepository
	duration time.Duration          // cache storage time
}

func NewUserCacheService(cache ports.ICacheRepository, duration time.Duration) ports.IUserCacheService {
	return &UserCacheService{
		cache,
		duration,
	}
}

// SetUser set an new user to cache
func (ucs *UserCacheService) SetUser(ctx context.Context, user *domain.User) error {
	bytes, err := marshal(user)
	if err != nil {
		return err
	}

	return ucs.cache.Set(ctx, fmt.Sprintf("%v-%v", userPrefix, user.ID), bytes, ucs.duration)
}

// GetUser get a user in cache by user id
func (ucs *UserCacheService) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	bytes, err := ucs.cache.Get(ctx, fmt.Sprintf("%v-%v", userPrefix, id))
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

// DeleteUser delete a user in cache
func (ucs *UserCacheService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ucs.cache.Delete(ctx, fmt.Sprintf("%v-%v", userPrefix, id))
}

// DeleteAllUsers delete all users in cache
func (ucs *UserCacheService) DeleteAllUsers(ctx context.Context) error {
	return ucs.cache.DeleteByPrefix(ctx, fmt.Sprintf("%v-*", userPrefix))
}
