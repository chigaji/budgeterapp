package controllers

import (
	"net/http"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

func GenerateFinancialReport(c echo.Context) error {

	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve user's expenses from the database
	var expenses []models.Expense
	models.DB.Where("user_id = ?", userID).Find(&expenses)

	// retrieve user's budgets from the database
	var budgets []models.Budget
	models.DB.Where("user_id = ?", userID).Find(&budgets)

	// calculate total income, total expenses and total budget
	var totalIncome float64
	var totalExpenses float64
	var totalBudget float64

	for _, expense := range expenses {
		totalExpenses += expense.Amount
	}

	for _, budget := range budgets {
		totalBudget += budget.Amount
	}

	// calculate total income by subtracting total expenses from total budget
	totalIncome = totalBudget - totalExpenses

	//generate report
	report := map[string]float64{
		"totalIncome":   totalIncome,
		"totalExpenses": totalExpenses,
		"totalBudget":   totalBudget,
	}

	// return report to the client
	return c.JSON(http.StatusOK, report)
}
