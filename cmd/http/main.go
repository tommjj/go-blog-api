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
	blogRepo := repository.NewBlogRepository(db)

	// cache
	userCache := cache.NewUserCache(redis, time.Hour)
	blogCache := cache.NewBlogCache(redis, time.Hour, time.Minute*2, time.Minute*2)

	// service
	tokenService, err := auth.NewJWTTokenService(*config.Auth)
	fatalOnError(err)

	authService := service.NewAuthService(tokenService, userRepo)
	userService := service.NewUserService(userRepo, userCache)
	blogService := service.NewBlogService(blogRepo, blogCache)

	// auth handler
	authHandler := handler.NewAuthHandler(authService)

	// user handler
	userHandler := handler.NewUserHandler(userService)

	// blog handler
	BlogHandler := handler.NewBlogHandler(blogService)

	r, err := http.New(config.Http,
		http.Group("/v1",
			http.RegisterAuthRoute(authHandler),
			http.RegisterUserRoute(tokenService, userHandler),
			http.RegisterBlogRoute(tokenService, BlogHandler),
		),
	)
	fatalOnError(err)

	r.Serve()
}
