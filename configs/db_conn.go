package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func loadEnv() {
	env := godotenv.Load()
	if env != nil {
		fmt.Printf("Error loading .env file: %v\n", env)
	}
}

func GetDBConnectionString() string {
	loadEnv()
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
}

func InitDB() *sql.DB {
	connStr := GetDBConnectionString()
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Successfully connected to the database!")
	return db
}
