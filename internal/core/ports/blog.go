package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type IBlogRepository interface {
	// GetBlogByID select a blog by id
	GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error)
	// GetBlogsByAuthorID select blogs by author id, with out blog text
	GetBlogsByAuthorID(ctx context.Context, id uuid.UUID, skip, limit int) ([]domain.Blog, error)
	// GetListBlogs get blogs
	GetListBlogs(ctx context.Context, skip, limit int) ([]domain.Blog, error)
	// SearchBlogsByName search blogs by name, with out blog text
	SearchBlogsByTitle(ctx context.Context, title string, skip, limit int) ([]domain.Blog, error)
	// CreateBlog insert an new blog into the database
	CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	// UpdateBlog update blog, only update non-zero fields by default
	UpdateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	// DeleteBlog delete blog by id
	DeleteBlog(ctx context.Context, id uuid.UUID) error
}

type IBlogCacheService interface {
	// SetBlog
	SetBlog(ctx context.Context, blog *domain.Blog) error
	// SetList
	SetList(ctx context.Context, skip int, limit int, list []domain.Blog) error
	// SetSearchList
	SetSearchList(ctx context.Context, search string, skip int, limit int, list []domain.Blog) error
	// GetBlog
	GetBlog(ctx context.Context, id uuid.UUID) (*domain.Blog, error)
	// GetList
	GetList(ctx context.Context, skip int, limit int) ([]domain.Blog, error)
	// GetSearchList
	GetSearchList(ctx context.Context, search string, skip int, limit int) ([]domain.Blog, error)
	// DeleteBlog
	DeleteBlog(ctx context.Context, id uuid.UUID) error
	// DeleteList
	DeleteList(ctx context.Context, skip int, limit int) error
	//DeleteSearchList
	DeleteSearchList(ctx context.Context, search string, skip int, limit int) error
	// DeleteSearchLists
	DeleteSearchLists(ctx context.Context, search string) error
	// DeleteAllList
	DeleteAllList(ctx context.Context) error
	// DeleteAllSearchList
	DeleteAllSearchList(ctx context.Context) error
	// DeleteAllBlogs
	DeleteAllBlogs(ctx context.Context) error
}

type IBlogService interface {
	// Authorized check if user owns blog
	Authorized(ctx context.Context, userId, blogId uuid.UUID) error
	GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error)
	GetListBlogs(ctx context.Context, skip, limit int) ([]domain.Blog, error)
	SearchBlogsByTitle(ctx context.Context, title string, skip, limit int) ([]domain.Blog, error)
	CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	UpdateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	DeleteBlog(ctx context.Context, blogId, userId uuid.UUID) error
}
