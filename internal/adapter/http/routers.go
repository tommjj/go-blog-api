package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/go-blog-api/internal/adapter/http/handler"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

// RegisterAuthRoute is a option function to return register auth router function
func RegisterAuthRoute(authHandler *handler.AuthHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		r := e.Group("/auth")
		{
			r.POST("/login", authHandler.Login)
		}
	}
}

// RegisterUserRoute is a option function to return register user router function
func RegisterUserRoute(token ports.ITokenService, authHandler *handler.UserHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		r := e.Group("/users")
		{
			r.GET("/:id", authHandler.GetUser)
			r.POST("/", authHandler.CreateUser)

			auth := r.Use(handler.AuthBeerMiddleware(token))
			{
				auth.PUT("/:id", authHandler.UpdateUser)
				auth.DELETE("/:id", authHandler.DeleteUser)
			}
		}
	}
}
