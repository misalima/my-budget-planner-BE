package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-budget-planner/cmd/app/handlers"
	"my-budget-planner/cmd/app/router"
	"my-budget-planner/internal/postgres"
	"my-budget-planner/internal/services"
	"os"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:*"}, // Allow any localhost port
		AllowCredentials: true,                           // Allows cookies and other credentials
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-CSRF-Token"},
	}))
	connStr, err := loadEnv()
	if err != nil {
		e.Logger.Fatal(err)
	}

	pool, err := postgres.ConnectDB(connStr)
	if err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Print("Connected to the database")
	}
	defer pool.Close()

	userServices := services.NewUserService(pool)
	userHandler := handlers.NewUserHandler(userServices)
	authServices := services.NewAuthService(pool)
	authHandler := handlers.NewAuthHandler(authServices)
	categoryServices := services.NewCategoryServices(pool)
	categoryHandler := handlers.NewCategoryHandler(categoryServices)
	creditCardServices := services.NewCreditCardService(pool)
	creditCardHandler := handlers.NewCreditCardHandler(creditCardServices)

	router.LoadRoutes(e, userHandler, authHandler, categoryHandler, creditCardHandler)
	e.Logger.Fatal(e.Start(":8000"))

}

func loadEnv() (string, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return "", errors.New("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s port=%s host=%s sslmode=disable",
		os.Getenv("MBP_PG_USER"),
		os.Getenv("MBP_PG_NAME"),
		os.Getenv("MBP_PG_PASSWORD"),
		os.Getenv("MBP_PG_PORT"),
		os.Getenv("MBP_PG_HOST"),
	)

	return connStr, nil
}
