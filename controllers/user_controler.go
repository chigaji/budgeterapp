package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	// get username and password from request body and store in database

	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	placeholderUser := User{
		Username: "chigaji",
		Phone:    "08012345678",
		Email:    "ronyg14@yahoo.com",
		Password: "test",
	}

	email := c.Param("email")
	password := c.Param("password")

	if password == placeholderUser.Password && email == placeholderUser.Email {
		log.Fatal("Login successful")
		return c.JSON(http.StatusOK, placeholderUser)
	}

	return c.JSON(http.StatusUnauthorized, "Login failed")

}
