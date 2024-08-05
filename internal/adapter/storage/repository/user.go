package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/adapter/storage"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/schema"
	"github.com/tommjj/go-blog-api/internal/core/domain"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *storage.DB
}

func NewUserRepository(db *storage.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.WithContext(ctx).Where("id = ?", userID).First(user).Error
	if err != nil {
		return nil, domain.ErrDataNotFound
	}

	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (ur *UserRepository) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.WithContext(ctx).Where("name = ?", name).First(user).Error
	if err != nil {
		return nil, domain.ErrDataNotFound
	}

	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// only update non-zero fields by default
func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userData := schema.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}

	newUserData := &schema.User{}

	upd := ur.db.WithContext(ctx).Clauses(clause.Returning{}).Model(newUserData).Where("id = ?", user.ID).Updates(userData)

	if err := upd.Error; err != nil {
		return nil, err
	}
	if row := upd.RowsAffected; row == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return &domain.User{
		ID:        newUserData.ID,
		Name:      newUserData.Name,
		Password:  newUserData.Password,
		CreatedAt: newUserData.CreatedAt,
		UpdatedAt: newUserData.UpdatedAt,
	}, nil
}

func (ur *UserRepository) UpdateUserByMap(ctx context.Context, id uuid.UUID, data *map[string]interface{}) (*domain.User, error) {
	newUserData := &schema.User{}

	upd := ur.db.WithContext(ctx).Clauses(clause.Returning{}).Model(newUserData).Omit("id").Where("id = ?", id).Updates(data)

	if err := upd.Error; err != nil {
		return nil, err
	}
	if row := upd.RowsAffected; row == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return &domain.User{
		ID:        newUserData.ID,
		Name:      newUserData.Name,
		Password:  newUserData.Password,
		CreatedAt: newUserData.CreatedAt,
		UpdatedAt: newUserData.UpdatedAt,
	}, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	var err error
	tx := ur.db.WithContext(ctx).Begin()

	if err = tx.Where("author_id = ?", id).Delete(&schema.Blog{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&schema.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
