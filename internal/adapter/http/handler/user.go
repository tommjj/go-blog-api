package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type UserHandler struct {
	us ports.IUserService
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		us: userService,
	}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required,min=3" example:"laplala" minLength:"3"`
	Password string `json:"password" binding:"required,min=8" example:"12345678" minLength:"8"`
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdUser, err := uh.us.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(createdUser)
	handleSuccess(ctx, res)
}

func (uh *UserHandler) GetUser(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	user, err := uh.us.GetUserByID(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(user)
	handleSuccess(ctx, res)
}

type updateUserRequest struct {
	Username string `json:"username" binding:"required,min=3" example:"laplala" minLength:"3"`
	Password string `json:"password" binding:"required,min=8" example:"12345678" minLength:"8"`
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest

	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	err = ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)
	if token.ID != id {
		handleError(ctx, domain.ErrForbidden)
		return
	}

	updateData := &domain.User{
		ID:       token.ID,
		Name:     req.Username,
		Password: req.Password,
	}

	updatedUser, err := uh.us.UpdateUser(ctx, updateData)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(updatedUser)
	handleSuccess(ctx, res)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)
	if token.ID != id {
		handleError(ctx, domain.ErrForbidden)
		return
	}

	err = uh.us.DeleteUser(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
