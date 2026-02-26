package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gilbert-keter/stellara/configs"
)

func main() {
	db := configs.InitDB()
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running")
	})

	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
