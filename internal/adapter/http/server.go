package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type RegisterRouterFunc func(gin.IRouter)

type Router struct {
	*gin.Engine
	Port int
	Url  string
}

func New(conf *config.Http, options ...RegisterRouterFunc) (*Router, error) {
	// gin.SetMode(gin.ReleaseMode)

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

	v1 := r.Group("/v1")
	for _, option := range options {
		option(v1)
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
