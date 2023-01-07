package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pavinjat/assessment/expenses"
)

func main() {

	expenses.InitDB()

	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" || password == "45678" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", expenses.CreateExpenseHandler)
	e.GET("/expenses/:id", expenses.GetExpenseHandler)

	log.Fatal(e.Start(":2565"))
}
