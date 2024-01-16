package controllers

import (
	"net/http"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

func AddExpense(c echo.Context) error {

	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	// pass and validate expense data from the request body

	var expense models.Expense

	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	expense.UserID = userID

	// add expense to database
	models.DB.Create(&expense)

	// return success response and status code to the client
	return c.JSON(http.StatusCreated, expense)
}

func GetExpenses(c echo.Context) error {

	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve expenses associated with the user from the database
	var expenses []models.Expense

	models.DB.Where("user_id = ?", userID).Find(&expenses)

	// return expenses to the client
	return c.JSON(http.StatusOK, expenses)
}

func UpdateExpense(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve and validate expense data from the request body

	var updateExpense models.Expense

	if err := c.Bind(&updateExpense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// make sure that the expense is associated with the authenticated user
	updateExpense.UserID = userID

	// update expense in the database
	models.DB.Save(&updateExpense)

	// return success response and status code to the client
	return c.JSON(http.StatusOK, updateExpense)
}

func DeleteExpense(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve expense id from the request param
	expenseID := c.Param("id")

	//delete expense from the database
	models.DB.Where("user_id = ? AND id = ?", userID, expenseID).Delete(&models.Expense{})

	// return success response and status code to the client
	// return c.JSON(http.StatusNoContent, map[string]string{"message": "Expense deleted successfully"})
	return c.NoContent(http.StatusNoContent)
}
