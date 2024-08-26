package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

type BlogHandler struct {
	svc ports.IBlogService
}

func NewBlogHandler(blogService ports.IBlogService) *BlogHandler {
	return &BlogHandler{
		svc: blogService,
	}
}

func (bh *BlogHandler) GetBlogByID(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	blog, err := bh.svc.GetBlogByID(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newBlogResponse(blog)
	handleSuccess(ctx, res)
}

type getListBlogsRequest struct {
	Query string `form:"q" binding:"" example:"how to ..."`
	Skip  int    `form:"skip" binding:"min=0" example:"0"`
	Limit int    `form:"limit" binding:"min=5" example:"5"`
}

func (bh *BlogHandler) GetListBlogs(ctx *gin.Context) {
	req := getListBlogsRequest{
		Limit: 5,
	}
	err := ctx.BindQuery(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	var blogs []domain.Blog
	if len(strings.TrimSpace(req.Query)) == 0 {
		blogs, err = bh.svc.GetListBlogs(ctx, req.Skip+1, req.Limit)
	} else {
		blogs, err = bh.svc.SearchBlogsByTitle(ctx, req.Query, req.Skip+1, req.Limit)
	}
	if err != nil {
		handleError(ctx, err)
		return
	}

	var res []blogResponse
	for _, blog := range blogs {
		res = append(res, newBlogResponse(&blog))
	}

	meta := newMeta(len(res), req.Skip, req.Limit)

	handleSuccess(ctx, responseWithMeta(meta, "blogs", res))
}

type createBlogRequest struct {
	Title string `json:"title" binding:"required" example:"adw..."`
	Text  string `json:"text" binding:"required" example:"adaw ..."`
}

func (bh *BlogHandler) CreateBlog(ctx *gin.Context) {
	var req createBlogRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	blog, err := bh.svc.CreateBlog(ctx, &domain.Blog{
		Title:    req.Title,
		Text:     req.Text,
		AuthorID: token.ID,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newBlogResponse(blog)
	handleSuccess(ctx, res)
}

type putBlogRequest struct {
	Title string `json:"title" binding:"required" example:"adw..."`
	Text  string `json:"text" binding:"required" example:"adaw ..."`
}

func (bh *BlogHandler) UpdateBlog(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	var req putBlogRequest
	err = ctx.BindJSON(&req)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	err = bh.svc.Authorized(ctx, token.ID, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	blog, err := bh.svc.UpdateBlog(ctx, &domain.Blog{
		ID:    id,
		Title: req.Title,
		Text:  req.Text,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	res := newBlogResponse(blog)
	handleSuccess(ctx, res)
}

func (bh *BlogHandler) DeleteBlog(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.Parse(paramId)
	if err != nil {
		validationError(ctx, err)
		return
	}

	token := getAuthPayload(ctx, authorizationPayloadKey)

	err = bh.svc.Authorized(ctx, token.ID, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	err = bh.svc.DeleteBlog(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
