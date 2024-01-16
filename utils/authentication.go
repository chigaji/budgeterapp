package utils

import (
	"github.com/chigaji/budgeterapp/models"
	"github.com/labstack/echo/v4"
)

func GenerateToken(user models.User) (string, error) {
	return "", nil
}

func ExtractUserIdFromToken(c echo.Context) (uint, error) {
	return 0, nil
}
