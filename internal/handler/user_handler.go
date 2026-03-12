package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gilbert-keter/stellara/internal/model"
	"github.com/gilbert-keter/stellara/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.Users = model.Users{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("content", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()

	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("content", "application/json")
	json.NewEncoder(w).Encode(users)
}
