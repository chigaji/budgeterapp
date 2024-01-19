package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

var blogger = utils.NewCustomLogger("utils/budget_controller")

func AddBudget(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	// pass and validate expense data from the request body

	var budget models.Budget

	if err := c.Bind(&budget); err != nil {
		blogger.Log(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	budget.UserID = userID

	// add expense to database
	models.DB.Create(&budget)

	blogger.Log(fmt.Sprint("Budget added : ", budget))
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

func GetBudget(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	//Get budget ID from req.param
	budgetID, err := strconv.Atoi(c.Param("id"))

	budget := models.Budget{}

	//retrieve the budget from the db
	if err := models.DB.First(&budget, "ID = ? AND user_id = ?", budgetID, userID).Error; err != nil {
		blogger.Log(fmt.Sprint("Error:", err))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}
	//return the budget
	return c.JSON(http.StatusOK, budget)
}

func UpdateBudget(c echo.Context) error {
	// extract user id from token
	userID, err := utils.ExtractUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	//Get budget ID from req.param
	budgetID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid budget ID"})
	}

	var updateBudget models.Budget

	// Retrieve the budget from the DB
	if err := models.DB.First(&updateBudget, budgetID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Budget not found"})
	}

	// retrieve and validate budget data from the request body
	if err := c.Bind(&updateBudget); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// ensure that expense is associated with an authenticated user
	updateBudget.UserID = userID

	// update expense in database
	models.DB.Save(&updateBudget)
	blogger.Log(fmt.Sprint("Updated Budget : ", updateBudget))
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
	logger1.Log(fmt.Sprint("Deleted Budget: ", budgetID))
	return c.NoContent(http.StatusNoContent)
}
