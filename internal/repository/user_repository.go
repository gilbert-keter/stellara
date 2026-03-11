package repository

import (
	"github.com/gilbert-keter/stellara/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user *model.Users) (model.Users, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return model.Users{}, err
	}
	return *user, nil
}

func (u *UserRepository) GetUsers() ([]model.Users, error) {
	var users []model.Users
	if err := u.DB.Find(&users).Error; err != nil {
		return []model.Users{}, err
	}
	return users, nil
}
