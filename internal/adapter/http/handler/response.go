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

// meta represents metadata for a paginated response
type meta struct {
	Total int `json:"total" example:"100"`
	Limit int `json:"limit" example:"10"`
	Skip  int `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip int) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
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

// blogResponse type to blog response for blog handler
type blogResponse struct {
	ID        uuid.UUID `json:"id" example:"39833b12-a044-46f5-8abd-47c47345d458"`
	Title     string    `json:"title" example:"how to ..."`
	Text      string    `json:"text,omitempty" example:"to do ..."`
	AuthorID  uuid.UUID `json:"author_id"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
}

// newBlogResponse create blog response for blog handler
func newBlogResponse(blog *domain.Blog) blogResponse {
	return blogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Text:      blog.Text,
		AuthorID:  blog.AuthorID,
		UpdatedAt: blog.UpdatedAt,
		CreatedAt: blog.CreatedAt,
	}
}

// listBlogsResponse type to blogs response for blog handler
type listBlogsResponse struct {
	Meta  meta           `json:"meta"`
	Blogs []blogResponse `json:"blogs"`
}

// newListBlogsResponse create blogs response for blog handler
func newListBlogsResponse(meta meta, blogs []blogResponse) listBlogsResponse {
	return listBlogsResponse{
		Meta:  meta,
		Blogs: blogs,
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
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"data not found"`
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
