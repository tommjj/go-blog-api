package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
	"github.com/tommjj/go-blog-api/internal/logger"
)

type BlogService struct {
	repo  ports.IBlogRepository
	cache ports.IBlogCacheService
}

func NewBlogService(blogRepository ports.IBlogRepository, cache ports.IBlogCacheService) ports.IBlogService {
	return &BlogService{
		repo:  blogRepository,
		cache: cache,
	}
}

func (bs *BlogService) GetBlogByID(ctx context.Context, id uuid.UUID) (*domain.Blog, error) {
	var blog *domain.Blog
	var err error

	blog, err = bs.cache.GetBlog(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			logger.Info(err.Error())
		} else {
			logger.Error(err.Error())
		}
	} else {
		return blog, nil
	}

	blog, err = bs.repo.GetBlogByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		} else {
			return nil, domain.ErrInternal
		}
	}

	err = bs.cache.SetBlog(ctx, blog)
	logIfErr(err)

	return blog, nil
}

func (bs *BlogService) GetListBlogs(ctx context.Context, skip, limit int) ([]domain.Blog, error) {
	var blogs []domain.Blog
	var err error

	blogs, err = bs.cache.GetList(ctx, skip, limit)
	if err != nil {
		if err == domain.ErrDataNotFound {
			logger.Info(err.Error())
		} else {
			logger.Error(err.Error())
		}
	} else {
		return blogs, nil
	}

	blogs, err = bs.repo.GetListBlogs(ctx, skip, limit)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		} else {
			return nil, domain.ErrInternal
		}
	}

	err = bs.cache.SetList(ctx, skip, limit, blogs)
	logIfErr(err)

	return blogs, nil
}

func (bs *BlogService) SearchBlogsByTitle(ctx context.Context, title string, skip, limit int) ([]domain.Blog, error) {
	var blogs []domain.Blog
	var err error

	blogs, err = bs.cache.GetSearchList(ctx, title, skip, limit)
	if err != nil {
		if err == domain.ErrDataNotFound {
			logger.Info(err.Error())
		} else {
			logger.Error(err.Error())
		}
	} else {
		return blogs, nil
	}

	blogs, err = bs.repo.SearchBlogsByTitle(ctx, title, skip, limit)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		} else {
			return nil, domain.ErrInternal
		}
	}

	err = bs.cache.SetSearchList(ctx, title, skip, limit, blogs)
	logIfErr(err)

	return blogs, nil
}

func (bs *BlogService) CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error) {
	newBlog, err := bs.repo.CreateBlog(ctx, blog)
	if err != nil {
		if err == domain.ErrDataConflict || err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	err = bs.cache.SetBlog(ctx, newBlog)
	logIfErr(err)

	return newBlog, nil
}

func (bs *BlogService) isAuthorized(ctx context.Context, blogId, userId uuid.UUID) error {
	blog, err := bs.GetBlogByID(ctx, blogId)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	if blog.AuthorID != userId {
		return domain.ErrUnauthorized
	}
	return nil
}

func (bs *BlogService) UpdateBlog(ctx context.Context, updates *domain.Blog) (*domain.Blog, error) {
	err := bs.isAuthorized(ctx, updates.ID, updates.AuthorID)
	if err != nil {
		return nil, err
	}

	updatedBlog, err := bs.repo.UpdateBlog(ctx, updates)
	if err != nil {
		if err == domain.ErrNoUpdatedData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	err = bs.cache.SetBlog(ctx, updatedBlog)
	logIfErr(err)
	return updatedBlog, nil
}

func (bs *BlogService) DeleteBlog(ctx context.Context, blogId, userId uuid.UUID) error {
	err := bs.isAuthorized(ctx, blogId, userId)
	if err != nil {
		return err
	}

	err = bs.repo.DeleteBlog(ctx, blogId)
	if err != nil {
		if err == domain.ErrNoUpdatedData {
			return err
		}
		return domain.ErrInternal
	}

	err = bs.cache.DeleteBlog(ctx, blogId)
	logIfErr(err)

	return nil
}
