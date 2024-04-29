package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/ramzyrsr/domain/user"
	"github.com/ramzyrsr/handlers"
	"github.com/ramzyrsr/infrastructure/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Reading environment variables
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbParams := os.Getenv("DB_PARAMS")

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	userRepository := postgres.NewUserRepositoryPG(db)

	// Initialize services
	userService := user.NewUserService(userRepository)

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	handlers.UserSetupRoutes(r, userService)

	// Run the server
	r.Run(":8080")
}
