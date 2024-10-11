package repository

import (
	"errors"

	"api-demo/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Database interface {
// 	Create(interface{}) *gorm.DB
// 	First(interface{}, ...interface{}) *gorm.DB
// }

// type UserRepository struct {
// 	Database
// }

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.Save(user).Error
}

func (r *UserRepository) GetUser(id uuid.UUID) (*model.User, error) {
	user := &model.User{ID: id}

	// SELECT * FROM users WHERE id = ?
	result := r.First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
