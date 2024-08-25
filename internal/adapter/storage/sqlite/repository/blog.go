package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite/schema"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
	"gorm.io/gorm/clause"
)

type BlogRepository struct {
	db *sqlite.DB
}

func NewBlogRepository(db *sqlite.DB) ports.IBlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (br *BlogRepository) GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error) {
	blog := &schema.Blog{}

	if err := br.db.WithContext(ctx).Where("id = ?", id).First(blog).Error; err != nil {
		return nil, domain.ErrDataNotFound
	}

	return &domain.Blog{
		ID:        blog.ID,
		Title:     blog.Title,
		Text:      blog.Text,
		AuthorID:  blog.AuthorID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}, nil
}

func (br *BlogRepository) GetBlogsByAuthorID(ctx context.Context, id uuid.UUID, skip, limit int) ([]domain.Blog, error) {
	blogs := []schema.Blog{}

	if err := br.db.WithContext(ctx).Select(
		"id", "title", "author_id", "created_at", "updated_at",
	).Where("author_id = ?", id).Limit(limit).Offset((skip - 1) * limit).Find(&blogs).Error; err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, domain.ErrDataNotFound
	}

	domainBlogs := []domain.Blog{}
	for _, blog := range blogs {
		domainBlogs = append(domainBlogs, domain.Blog{
			ID:        blog.ID,
			Title:     blog.Title,
			Text:      blog.Text,
			AuthorID:  blog.AuthorID,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}
	return domainBlogs, nil
}

func (br *BlogRepository) GetListBlogs(ctx context.Context, skip, limit int) ([]domain.Blog, error) {
	blogs := []schema.Blog{}

	err := br.db.WithContext(ctx).Select(
		"id", "title", "author_id", "created_at", "updated_at",
	).Limit(limit).Offset((skip - 1) * limit).Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, domain.ErrDataNotFound
	}

	domainBlogs := []domain.Blog{}
	for _, blog := range blogs {
		domainBlogs = append(domainBlogs, domain.Blog{
			ID:        blog.ID,
			Title:     blog.Title,
			Text:      blog.Text,
			AuthorID:  blog.AuthorID,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}
	return domainBlogs, nil
}

func (br *BlogRepository) SearchBlogsByTitle(ctx context.Context, title string, skip, limit int) ([]domain.Blog, error) {
	blogs := []schema.Blog{}

	err := br.db.WithContext(ctx).Select(
		"id", "title", "author_id", "created_at", "updated_at",
	).Where("title LIKE ?", fmt.Sprintf("%%%v%%", title)).Limit(limit).Offset((skip - 1) * limit).Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, domain.ErrDataNotFound
	}

	domainBlogs := []domain.Blog{}
	for _, blog := range blogs {
		domainBlogs = append(domainBlogs, domain.Blog{
			ID:        blog.ID,
			Title:     blog.Title,
			Text:      blog.Text,
			AuthorID:  blog.AuthorID,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}
	return domainBlogs, nil
}

func (br *BlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error) {
	// sqlite is not support check forget key
	if err := br.db.Where("id = ?", blog.AuthorID).First(&schema.User{}).Error; err != nil {
		return nil, domain.ErrDataConflict
	}

	newBlog := &schema.Blog{
		Title:    blog.Title,
		Text:     blog.Text,
		AuthorID: blog.AuthorID,
	}

	if err := br.db.WithContext(ctx).Create(newBlog).Error; err != nil {
		return nil, err
	}

	return &domain.Blog{
		ID:        newBlog.ID,
		Title:     newBlog.Title,
		Text:      newBlog.Text,
		AuthorID:  newBlog.AuthorID,
		CreatedAt: newBlog.CreatedAt,
		UpdatedAt: newBlog.UpdatedAt,
	}, nil
}

func (br *BlogRepository) UpdateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error) {
	updateData := &schema.Blog{
		Title: blog.Title,
		Text:  blog.Text,
	}
	updatedData := &schema.Blog{}

	upd := br.db.WithContext(ctx).Clauses(clause.Returning{}).Model(updatedData).Where("id = ?", blog.ID).Updates(updateData)
	if err := upd.Error; err != nil {
		return nil, err
	}
	if row := upd.RowsAffected; row == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return &domain.Blog{
		ID:        updatedData.ID,
		Title:     updatedData.Title,
		Text:      updatedData.Text,
		AuthorID:  updatedData.AuthorID,
		CreatedAt: updatedData.CreatedAt,
		UpdatedAt: updatedData.UpdatedAt,
	}, nil
}

func (br *BlogRepository) DeleteBlog(ctx context.Context, id uuid.UUID) error {
	dl := br.db.WithContext(ctx).Delete(&schema.Blog{}, id)

	if err := dl.Error; err != nil {
		return err
	}
	if dl.RowsAffected == 0 {
		return domain.ErrNoUpdatedData
	}

	return nil
}
