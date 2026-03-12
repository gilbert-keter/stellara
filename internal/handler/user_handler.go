package handler

import (
	"encoding/json"
	"fmt"
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
	defer r.Body.Close()

	var user model.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()

	if err != nil {
		fmt.Errorf("Error: %w\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
