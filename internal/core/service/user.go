package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"github.com/tommjj/go-blog-api/internal/core/ports"
	"github.com/tommjj/go-blog-api/internal/core/util"
	"github.com/tommjj/go-blog-api/internal/logger"
)

type UserService struct {
	repo  ports.IUserRepository   // user repo
	cache ports.IUserCacheService // user cache
}

func NewUserService(userRepo ports.IUserRepository, cache ports.IUserCacheService) ports.IUserService {
	return &UserService{
		repo:  userRepo,
		cache: cache,
	}
}

func (us *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user *domain.User
	var err error

	user, err = us.cache.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			logger.Info(err.Error())
		} else {
			logger.Error(err.Error())
		}
	} else {
		return user, nil
	}

	user, err = us.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Password = "" // remove password

	err = us.cache.SetUser(ctx, user)
	logOnError(err)

	return user, nil
}

func (us *UserService) CreateUser(ctx context.Context, username, password string) (*domain.User, error) {
	hashPass, err := util.HashPassword(password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	user, err := us.repo.CreateUser(ctx, &domain.User{
		Name:     username,
		Password: hashPass,
	})
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	user.Password = ""

	err = us.cache.SetUser(ctx, user)
	logOnError(err)

	return user, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Name == "" && user.Password == "" {
		return nil, domain.ErrNoUpdatedData
	}

	existingUser, err := us.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if user.Password == "" && user.Name == existingUser.Name {
		return nil, domain.ErrNoUpdatedData
	}

	hashPass := ""
	if user.Password != "" {
		hashPass, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	updatedUser, err := us.repo.UpdateUser(ctx, &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: hashPass,
	})
	if err != nil {
		if err == domain.ErrConflictingData || err == domain.ErrNoUpdatedData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	updatedUser.Password = ""

	err = us.cache.DeleteUser(ctx, user.ID)
	logOnError(err)

	err = us.cache.SetUser(ctx, user)
	logOnError(err)

	return updatedUser, nil
}

func (us *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := us.repo.DeleteUser(ctx, id)

	if err != nil {
		if err == domain.ErrNoUpdatedData {
			return err
		}
		return domain.ErrInternal
	}

	err = us.cache.DeleteUser(ctx, id)
	logOnError(err)

	return nil
}
