package controllers

import (
	"log"
	"net/http"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

func AddBudget(c echo.Context) error {

	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	log.Fatal(userID)

	//add budget to database and return to client

	return nil
}

func getBudgets(c echo.Context) error {
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	log.Fatal(userID)

	// get all budgets from database and return to client
	return nil
}

func AddBugdget(c echo.Context) error {

	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	// pass and validate expense data from the request body

	var budget models.Budget

	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	budget.UserID = userID

	// add expense to database
	models.DB.Create(&budget)

	// return success response and status code to the client
	return c.JSON(http.StatusCreated, budget)
}

func GetBudgets(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve expenses associated with the user from the database
	var budgets []models.Budget

	models.DB.Where("user_id = ?", userID).Find(&budgets)

	// return expenses to the client
	return c.JSON(http.StatusOK, budgets)
}

func UpdateBudget(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// retrieve and validate expense data from the request body

	var updateBudget models.Budget

	if err := c.Bind(&updateBudget); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	updateBudget.UserID = userID

	// update expense in database
	models.DB.Save(&updateBudget)

	// return success response and status code to the client
	return c.JSON(http.StatusOK, updateBudget)
}

func DeleteBudget(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	// pass and validate budget Id from the request

	budgetID := c.Param("id")

	//delete budget from database
	models.DB.Where("id = ? AND user_id = ?", budgetID, userID).Delete(&models.Budget{})

	// return success response and status code to the client
	// return c.JSON(http.StatusOK, map[string]string{"message": "budget deleted successfully"})
	return c.NoContent(http.StatusNoContent)
}
