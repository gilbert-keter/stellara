package routes

import (
	"net/http"

	"github.com/gilbert-keter/stellara/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, taskHandler *handler.TaskHandler) http.Handler {
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/health", handler.HealthHandler).Methods("GET")
	r.HandleFunc("/task/add", taskHandler.CreateTask).Methods("POST")
	return r
}

func RegisterUserRoutes(r *mux.Router, userHandler *handler.UserHandler) http.Handler {
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/user/add", userHandler.CreateUser).Methods("POST")
	return r
}
