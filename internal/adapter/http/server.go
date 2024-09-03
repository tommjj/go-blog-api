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

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/tommjj/go-blog-api/docs"
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
	if conf.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// set logger middleware
	logger, err := logger.New(conf.Logger)
	if err != nil {
		return nil, errors.New("http logger conf is not valid")
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// set CORS
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = conf.AllowedOrigins
	r.Use(cors.New(ginConfig))

	// set router
	r.GET("/ping", ping)

	for _, option := range options {
		option(r)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/v1/api"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{
		Engine: r,
		Port:   conf.Port,
		Url:    conf.URL,
	}, nil
}

func (r *Router) Serve() {
	logger.Info(fmt.Sprintf("start server at http://%v:%v", r.Url, r.Port))

	err := r.Run(fmt.Sprintf("%v:%v", r.Url, r.Port))
	if err != nil {
		logger.Error(err.Error())
	}
}
