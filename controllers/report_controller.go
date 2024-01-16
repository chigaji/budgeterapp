package controllers

import (
	"log"
	"net/http"

	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

func GenerateFinancialReport(c echo.Context) error {

	userID, err := utils.ExtractUserIdFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	log.Fatal(userID)

	log.Fatal("Generating financial report")
	// generate financial report and return to client

	return nil
}
