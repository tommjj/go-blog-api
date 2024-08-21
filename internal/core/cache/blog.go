package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
)

var (
	blogPrefix            = "blog"
	listBlogsPrefix       = "blogs"
	searchListBlogsPrefix = "searchBlogs"
)

type BlogCacheService struct {
	cache          ports.ICacheRepository
	blogDuration   time.Duration
	listDuration   time.Duration
	searchDuration time.Duration
}

func NewBlogCacheService(cache ports.ICacheRepository, blogDuration time.Duration, listDuration time.Duration, searchDuration time.Duration) ports.IBlogCacheService {
	return &BlogCacheService{
		cache:          cache,
		blogDuration:   blogDuration,
		listDuration:   listDuration,
		searchDuration: searchDuration,
	}
}

func (bcs *BlogCacheService) SetBlog(ctx context.Context, blog *domain.Blog) error {
	bytes, err := marshal(blog)
	if err != nil {
		return err
	}

	return bcs.cache.Set(ctx, generateCacheKeyParams(blogPrefix, blog.ID), bytes, bcs.blogDuration)
}

func (bcs *BlogCacheService) SetList(ctx context.Context, skip int, limit int, list []domain.Blog) error {
	bytes, err := marshal(list)
	if err != nil {
		return err
	}

	return bcs.cache.Set(ctx, generateCacheKeyParams(listBlogsPrefix, skip, limit), bytes, bcs.listDuration)
}

func (bcs *BlogCacheService) SetSearchList(ctx context.Context, search string, skip int, limit int, list []domain.Blog) error {
	bytes, err := marshal(list)
	if err != nil {
		return err
	}

	return bcs.cache.Set(ctx, generateCacheKeyParams(searchListBlogsPrefix, search, skip, limit), bytes, bcs.searchDuration)
}

func (bcs *BlogCacheService) GetBlog(ctx context.Context, id uuid.UUID) (*domain.Blog, error) {
	bytes, err := bcs.cache.Get(ctx, generateCacheKeyParams(blogPrefix, id))
	if err != nil {
		return nil, err
	}

	blog := &domain.Blog{}
	err = unmarshal(bytes, blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (bcs *BlogCacheService) GetList(ctx context.Context, skip int, limit int) ([]domain.Blog, error) {
	bytes, err := bcs.cache.Get(ctx, generateCacheKeyParams(listBlogsPrefix, skip, limit))
	if err != nil {
		return nil, err
	}

	list := []domain.Blog{}
	err = unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (bcs *BlogCacheService) GetSearchList(ctx context.Context, search string, skip int, limit int) ([]domain.Blog, error) {
	bytes, err := bcs.cache.Get(ctx, generateCacheKeyParams(searchListBlogsPrefix, search, skip, limit))
	if err != nil {
		return nil, err
	}

	list := []domain.Blog{}
	err = unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (bcs *BlogCacheService) DeleteBlog(ctx context.Context, id uuid.UUID) error {
	return bcs.cache.Delete(ctx, generateCacheKeyParams(blogPrefix, id))
}

func (bcs *BlogCacheService) DeleteList(ctx context.Context, skip int, limit int) error {
	return bcs.cache.Delete(ctx, generateCacheKeyParams(listBlogsPrefix, skip, limit))
}

func (bcs *BlogCacheService) DeleteSearchList(ctx context.Context, search string, skip int, limit int) error {
	return bcs.cache.Delete(ctx, generateCacheKeyParams(searchListBlogsPrefix, search, skip, limit))
}

func (bcs *BlogCacheService) DeleteSearchLists(ctx context.Context, search string) error {
	return bcs.cache.DeleteByPrefix(ctx, fmt.Sprintf("%v-%v*", searchListBlogsPrefix, search))
}

func (bcs *BlogCacheService) DeleteAllList(ctx context.Context) error {
	return bcs.cache.DeleteByPrefix(ctx, fmt.Sprintf("%v-*", listBlogsPrefix))
}

func (bcs *BlogCacheService) DeleteAllSearchList(ctx context.Context) error {
	return bcs.cache.DeleteByPrefix(ctx, fmt.Sprintf("%v-*", searchListBlogsPrefix))
}

func (bcs *BlogCacheService) DeleteAllBlogs(ctx context.Context) error {
	return bcs.cache.DeleteByPrefix(ctx, fmt.Sprintf("%v-*", blogPrefix))
}
