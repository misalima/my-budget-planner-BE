package main

import (
	"github.com/labstack/echo/v4"
	"my-budget-planner/cmd/app/router"
)

func main() {
	e := echo.New()
	router.LoadRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))

}
