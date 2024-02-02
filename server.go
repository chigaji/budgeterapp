package main

import (
	"net/http"

	"github.com/chigaji/budgeterapp/controllers"
	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// connect to database
	models.ConnectToDatabase()

	e := echo.New()

	//add cors
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://localhost:3000"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// }))
	// Use custom logger
	customLogger := utils.NewCustomLogger("main")

	//customize logging

	// Middleware
	e.Use(utils.NewCustomHTTPLogger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		customLogger.Log("Welcome to the financial budgeter app's API!")
		return c.JSON(http.StatusOK, "Welcome to the financial budgeter app's API!")
	})

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Restricted routes
	r := e.Group("/api/v1")

	r.Use(middleware.JWT([]byte(utils.JwtSecrete)))

	r.GET("/expenses", controllers.GetExpenses)
	r.POST("/expenses", controllers.AddExpense)
	r.GET("/expenses/:id", controllers.GetExpense)
	r.PUT("/expenses/:id", controllers.UpdateExpense)
	r.DELETE("/expenses/:id", controllers.DeleteExpense)

	r.POST("/budgets", controllers.AddBudget)
	r.GET("/budgets", controllers.GetBudgets)
	r.GET("/budgets/:id", controllers.GetBudget)
	r.PUT("/budgets/:id", controllers.UpdateBudget)
	r.DELETE("/budgets/:id", controllers.DeleteBudget)

	r.GET("/reports", controllers.GenerateFinancialReport)

	e.Logger.Fatal(e.Start(":1323"))

}
