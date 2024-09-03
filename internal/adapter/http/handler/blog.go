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

// GetBlog go-blog
//
//	@Summary		get blog
//	@Description	get blog by blog id
//	@Tags			blogs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"blog id"	format(uuid)
//	@Success		200	{object}	response{data=blogResponse}	"Blog data"
//	@Failure		400	{object}	errorResponse				"Validation error"
//	@Failure		404	{object}	errorResponse				"Data not found error"
//	@Failure		500	{object}	errorResponse				"Internal server error"
//	@Router			/blogs/{id} [get]
func (bh *BlogHandler) GetBlog(ctx *gin.Context) {
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

// GetListBlogs go-blog
//
//	@Summary		get blogs
//	@Description	get blogs
//	@Tags			blogs
//	@Accept			json
//	@Produce		json
//	@Param			q		query		string								false	"Query"
//	@Param			skip	query		int									false	"Skip"	default(0)	minimum(0)
//	@Param			limit	query		int									false	"Limit"	default(5)	minimum(5)
//	@Success		200		{object}	response{data=listBlogsResponse}	"Blogs data"
//	@Failure		400		{object}	errorResponse						"Validation error"
//	@Failure		404		{object}	errorResponse						"Data not found error"
//	@Failure		500		{object}	errorResponse						"Internal server error"
//	@Router			/blogs [get]
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

	res := make([]blogResponse, 0, len(blogs))
	for _, blog := range blogs {
		res = append(res, newBlogResponse(&blog))
	}

	meta := newMeta(len(res), req.Limit, req.Skip)

	handleSuccess(ctx, newListBlogsResponse(meta, res))
}

type createBlogRequest struct {
	Title string `json:"title" binding:"required" example:"adw..."`
	Text  string `json:"text" binding:"required" example:"adaw ..."`
}

// CreateBlog go-blog
//
//	@Summary		create blog
//	@Description	create a new blog
//	@Tags			blogs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createBlogRequest			true	"Create blog request body"
//	@Success		200		{object}	response{data=blogResponse}	"Blog created"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		409		{object}	errorResponse				"Data conflict error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/blogs [post]
//	@Security		BearerAuth
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

// CreateBlog go-blog
//
//	@Summary		update blog
//	@Description	update a blog data
//	@Tags			blogs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Blog id"	format(uuid)
//	@Param			request	body		putBlogRequest				true	"Update blog request body"
//	@Success		200		{object}	response{data=blogResponse}	"Blog updated"
//	@Failure		400		{object}	errorResponse				"Validation error"
//	@Failure		401		{object}	errorResponse				"Unauthorized error"
//	@Failure		403		{object}	errorResponse				"Forbidden error"
//	@Failure		409		{object}	errorResponse				"Data conflict error"
//	@Failure		500		{object}	errorResponse				"Internal server error"
//	@Router			/blogs/{id} [put]
//	@Security		BearerAuth
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

// DeleteBlog go-blog
//
//	@Summary		delete blog
//	@Description	delete a blog
//	@Tags			blogs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Blog id"	format(uuid)
//	@Success		200	{object}	response		"Blog updated"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		409	{object}	errorResponse	"Data conflict error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/blogs/{id} [delete]
//	@Security		BearerAuth
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
