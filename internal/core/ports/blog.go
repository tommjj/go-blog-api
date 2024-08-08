package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type IBlogRepository interface {
	// GetBlogByID select a blog by id
	GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error)
	// GetBlogsByAuthorID select blogs by author id
	GetBlogsByAuthorID(ctx context.Context, id uuid.UUID, skip, limit int) ([]domain.Blog, error)
	// SearchBlogsByName search blogs by name
	SearchBlogsByName(ctx context.Context, name string, skip, limit int) ([]domain.Blog, error)
	// CreateBlog insert an new blog into the database
	CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	// UpdateBlog update blog, only update non-zero fields by default
	UpdateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	// DeleteBlog delete blog by id
	DeleteBlog(ctx context.Context, id uuid.UUID) error
}
