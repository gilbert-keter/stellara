package service

import (
	"github.com/gilbert-keter/stellara/internal/model"
	"github.com/gilbert-keter/stellara/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserSevice(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}

}

func (s *UserService) CreateUser(user *model.Users) (model.Users, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]model.Users, error) {
	return s.repo.GetUsers()
}
