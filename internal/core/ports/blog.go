package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
)

type IBlogRepository interface {
	GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error)
	GetListBlogByAuthorID(ctx context.Context, id uuid.UUID, skip, limit int) ([]domain.Blog, error)
	SearchListBlogByName(ctx context.Context, name string, skip, limit int) ([]domain.Blog, error)
	CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	UpdateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error)
	DeleteBlog(ctx context.Context, id uuid.UUID) error
}
