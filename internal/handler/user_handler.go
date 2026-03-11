package handler

import (
	"github.com/gilbert-keter/stellara/internal/model"
	"github.com/gilbert-keter/stellara/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func (h *UserHandler) CreateUser(user *model.Users) (model.Users, error) {
	return h.service.CreateUser(user)
}

func (h *UserHandler) GetUsers() ([]model.Users, error) {
	return h.service.GetUsers()
}
