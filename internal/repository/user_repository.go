package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"api-demo/internal/model"
)

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

	result := r.First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
