package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

var logger1 = utils.NewCustomLogger("controllers/expense_controller")

func AddExpense(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	logger1.Log(fmt.Sprint("User Id : ", userID))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	// pass and validate expense data from the request body
	var expense models.Expense

	if err := c.Bind(&expense); err != nil {
		logger1.Log(fmt.Sprint("Error: ", err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	expense.UserID = userID

	// add expense to database
	models.DB.Create(&expense)

	// log output
	logger1.Log(fmt.Sprint("Expense created: ", expense))

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

func GetExpense(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	//Get expense ID from req.param
	expenseID, err := strconv.Atoi(c.Param("id"))

	expense := models.Expense{}

	//retrieve the expense from the db
	if err := models.DB.First(&expense, "ID = ? AND user_id = ?", expenseID, userID).Error; err != nil {
		blogger.Log(fmt.Sprint("Error:", err))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}
	//return the expense
	return c.JSON(http.StatusOK, expense)
}

func UpdateExpense(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	//Get expense ID from req.param
	expenseID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expense ID"})
	}

	var updateExpense models.Expense

	// Retrieve the expense from the DB
	if err := models.DB.First(&updateExpense, expenseID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Expense not found"})
	}

	// retrieve and validate expense data from the request body
	if err := c.Bind(&updateExpense); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// make sure that the expense is associated with the authenticated user
	updateExpense.UserID = userID

	// update expense in the database
	models.DB.Save(&updateExpense)

	logger1.Log(fmt.Sprint("Updated Expense : ", updateExpense))

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

	logger1.Log(fmt.Sprint("Deleted expense: ", expenseID))
	return c.NoContent(http.StatusNoContent)
}
