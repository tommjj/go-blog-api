package service

import (
	"time"

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

// func (ucs *UserCacheService) Set
