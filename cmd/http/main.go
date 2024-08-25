package main

import (
	"context"
	"time"

	"github.com/tommjj/go-blog-api/internal/adapter/http"
	"github.com/tommjj/go-blog-api/internal/adapter/http/handler"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/redis"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite/repository"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/core/auth"
	"github.com/tommjj/go-blog-api/internal/core/cache"
	"github.com/tommjj/go-blog-api/internal/core/service"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func fatalOnError(err error) {
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func main() {
	config, err := config.New()
	fatalOnError(err)

	// setup logger
	err = logger.Set(*config.Logger)
	fatalOnError(err)
	defer logger.Sync()

	// database
	db, err := sqlite.New(*config.DB)
	fatalOnError(err)

	//redis
	redis, err := redis.New(context.Background(), *config.Redis)
	fatalOnError(err)
	defer redis.Close()

	// repository
	userRepo := repository.NewUserRepository(db)

	// cache
	userCache := cache.NewUserCache(redis, time.Hour)

	// auth
	tokenService, err := auth.NewJWTTokenService(*config.Auth)
	fatalOnError(err)
	authService := service.NewAuthService(tokenService, userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// user handler
	userService := service.NewUserService(userRepo, userCache)
	userHandler := handler.NewUserHandler(userService)

	r, err := http.New(config.Http,
		http.Group("/v1",
			http.RegisterAuthRoute(authHandler),
			http.RegisterUserRoute(tokenService, userHandler),
		),
	)
	fatalOnError(err)

	r.Serve()
}
