package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-budget-planner/cmd/app/handlers"
	"my-budget-planner/cmd/app/router"
	"my-budget-planner/internal/postgres"
	"my-budget-planner/internal/services"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:*"}, // Allow any localhost port
		AllowCredentials: true,                           // Allows cookies and other credentials
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-CSRF-Token"},
	}))
	connStr := "user=postgres dbname=mbp_pg_db password=12345678 port=5432 host=localhost sslmode=disable"
	pool, err := postgres.ConnectDB(connStr)
	if err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Print("Connected to the database")
	}
	defer pool.Close()

	userServices := services.NewUserService(pool)
	userHandler := handlers.NewUserHandler(userServices)

	router.LoadRoutes(e, userHandler)
	e.Logger.Fatal(e.Start(":8000"))

}
