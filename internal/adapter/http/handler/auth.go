package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type AuthHandler struct {
	as ports.IAuthService
}

func NewAuthHandler(authService ports.IAuthService) *AuthHandler {
	return &AuthHandler{
		as: authService,
	}
}

type loginRequest struct {
	Username string `json:"username" binding:"required,min=3" example:"laplala" minLength:"3"`
	Password string `json:"password" binding:"required,min=8" example:"12345678" minLength:"8"`
}

func (auth AuthHandler) Login(ctx *gin.Context) {
	var req loginRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token, err := auth.as.Login(ctx, req.Username, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}
	res := newAuthResponse(token)

	handleSuccess(ctx, res)
}
