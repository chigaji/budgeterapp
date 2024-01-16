package controllers

import (
	"net/http"

	"github.com/chigaji/budgeterapp/models"
	"github.com/chigaji/budgeterapp/utils"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// hash user password
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error hashing password"})
	}

	user.Password = hashedPassword

	// add user to database
	models.DB.Create(&user)

	// return success response and status code to the client
	return c.JSON(http.StatusCreated, user)

}

func Login(c echo.Context) error {

	// pass and validate user login data from the request body
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// retrieve user from database using username
	var user models.User

	if err := models.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// verify creddentials and generate token

	if err := utils.ValidatePassword(user.Password, loginData.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating token"})
	}

	// return token to the client
	return c.JSON(http.StatusOK, map[string]string{"token": token})

}
