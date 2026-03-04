package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gilbert-keter/stellara/configs"
	"github.com/gilbert-keter/stellara/internal/handler"
	"github.com/gilbert-keter/stellara/internal/model"
	"github.com/gilbert-keter/stellara/internal/repository"
	"github.com/gilbert-keter/stellara/internal/routes"
	"github.com/gilbert-keter/stellara/internal/service"
)

func main() {
	// Initialize GORM DB
	db := configs.InitGormDB()

	// Auto-migrate first
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Initialize repository, service, handler
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Register routes (pass handlers to routes if needed)
	router := routes.RegisterRoutes(taskHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server failed:", err)
	}
}