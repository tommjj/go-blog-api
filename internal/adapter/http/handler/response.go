package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// authResponse type to auth response for auth handler
type authResponse struct {
	Token string `json:"token" example:"eyJJ9.eyJpEzNDR9.fUjDw0"`
}

// newAuthResponse create a auth response for login handler
func newAuthResponse(token string) authResponse {
	return authResponse{
		Token: token,
	}
}

// userResponse type to user response for user handler
type userResponse struct {
	ID        uuid.UUID `json:"id" example:"39833b12-a044-46f5-8abd-47c47345d458"`
	Username  string    `json:"username" example:"laplala"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
}

// newUserResponse create user response for user handler
func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Name,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
}

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
}

// handleSuccess write success response with status code 200 mess Success and data
func handleSuccess(ctx *gin.Context, data any) {
	res := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, res)
}

// handleError write error response
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := praseError(err)
	res := newErrorResponse(errMsg)
	ctx.JSON(statusCode, res)
}

// errorResponse type of error response
type errorResponse struct {
	Success  bool     `json:"success" example:"true"`
	Messages []string `json:"messages" example:"Success"`
}

// newErrorResponse create an new error response
func newErrorResponse(errMegs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMegs,
	}
}

// validationError handle validation error, write err response
func validationError(ctx *gin.Context, err error) {
	errMegs := praseError(err)
	res := newErrorResponse(errMegs)

	ctx.JSON(http.StatusBadRequest, res)
}

// praseError prase error to error messages
func praseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, validationError := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, validationError.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}
	return errMsgs
}
