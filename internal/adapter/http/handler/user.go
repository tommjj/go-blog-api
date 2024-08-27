package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type UserHandler struct {
	svc ports.IUserService
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		svc: userService,
	}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required,min=3" example:"laplala" minLength:"3"`
	Password string `json:"password" binding:"required,min=8" example:"12345678" minLength:"8"`
}

// CreateUser go-blog
//
//	@Summary		create user
//	@Description	create an new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createUserRequest			true	"Create User request body"
//	@Success		200		{object}	response{data=userResponse}	"User created"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		409		{object}	errorResponse				"Data conflict error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/users [post]
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	createdUser, err := uh.svc.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(createdUser)
	handleSuccess(ctx, res)
}

// GetUser go-blog
//
//	@Summary		get user
//	@Description	get a user by user id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uuid						true	"User id"
//	@Success		200	{object}	response{data=userResponse}	"User data"
//	@Failure		400	{object}	errorResponse				"Validation error"
//	@Failure		404	{object}	errorResponse				"Data not found error"
//	@Failure		500	{object}	errorResponse				"Internal server error"
//	@Router			/users/{id} [get]
func (uh *UserHandler) GetUser(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	user, err := uh.svc.GetUserByID(ctx, id)
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

// UpdateUser go-blog
//
//	@Summary		update user
//	@Description	update user data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uuid						true	"User id"
//	@Param			request	body		updateUserRequest			true	"Update User request body"
//	@Success		200		{object}	response{data=userResponse}	"User updated"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		409		{object}	errorResponse				"Data conflict error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
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

	updatedUser, err := uh.svc.UpdateUser(ctx, updateData)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newUserResponse(updatedUser)
	handleSuccess(ctx, res)
}

// DeleteUser go-blog
//
//	@Summary		delete user
//	@Description	delete user by user id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uuid			true	"User id"
//	@Success		200	{object}	response		"User deleted"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
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

	err = uh.svc.DeleteUser(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
