package main

import (
	"fmt"
	"net/http"

	"github.com/gilbert-keter/stellara/configs"
	"github.com/gilbert-keter/stellara/internal/routes"
)

func main() {
	db := configs.InitDB()
	defer db.Close()
	router := routes.RegisterRoutes()
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, router)
}
