package routes

import (
	"net/http"

	"github.com/gilbert-keter/stellara/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes(taskHandler *handler.TaskHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/health", handler.HealthHandler).Methods("GET")
	r.HandleFunc("/task/add", taskHandler.CreateTask).Methods("POST")
	return r
}
