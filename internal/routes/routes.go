package routes

import (
	"net/http"

	"github.com/gilbert-keter/stellara/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/health", handler.HealthHandler).Methods("GET")
	return r
}
