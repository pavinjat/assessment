package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pavinjat/assessment/config"
	"github.com/pavinjat/assessment/expenses"
)

func main() {

	config := config.NewConfig()
	expenses.InitDB()

	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "expensesapi" || password == "123456" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", expenses.CreateExpenseHandler)
	e.GET("/expenses", expenses.GetAllExpensesHandler)
	e.GET("/expenses/:id", expenses.GetExpenseHandler)
	e.PUT("/expenses/:id", expenses.UpdateExpenseHandler)

	go func() {
		if err := e.Start(config.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
