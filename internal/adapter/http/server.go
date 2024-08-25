package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommjj/go-blog-api/internal/adapter/http/handler"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type Router struct {
	*gin.Engine
	Port int
	Url  string
}

func New(conf *config.Http, authHandler *handler.AuthHandler) (*Router, error) {
	r := gin.New()

	// set logger middleware
	logger, err := logger.New(conf.Logger)
	if err != nil {
		return nil, errors.New("http logger conf is not valid")
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// CORS
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = conf.AllowedOrigins
	r.Use(cors.New(ginConfig))

	// router
	r.GET("/ping", ping)

	//auth
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	return &Router{
		Engine: r,
		Port:   conf.Port,
		Url:    conf.URL,
	}, nil
}

func (r *Router) Serve() {
	r.Run(fmt.Sprintf("%v:%v", r.Url, r.Port))
}
