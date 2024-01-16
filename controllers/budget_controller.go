package controllers

import (
	"log"
	"net/http"

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