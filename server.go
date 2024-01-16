package main

import (
	"net/http"

	"github.com/chigaji/budgeterapp/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to the financial budgeter app's API!")
	})

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	r := e.Group("/api/v1")
	r.Use((middleware.JWT([]byte("secret"))))

	r.GET("/expenses", controllers.GetExpenses)
	r.POST("/expenses", controllers.AddExpense)
	r.PUT("/expenses/:id", controllers.UpdateExpense)
	r.DELETE("/expenses/:id", controllers.DeleteExpense)

	r.GET("/budgets", controllers.GetBudgets)
	r.POST("/budgets", controllers.AddBudget)
	r.PUT("/budgets/:id", controllers.UpdateBudget)
	r.DELETE("/budgets/:id", controllers.DeleteBudget)

	r.GET("/reports", controllers.GenerateFinancialReport)

	e.Logger.Fatal(e.Start(":1323"))

}
