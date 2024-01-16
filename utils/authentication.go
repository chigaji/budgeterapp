package utils

import (
	"time"

	"github.com/chigaji/budgeterapp/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const JwtSecrete = "budgeterapp"

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // token expires after 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(JwtSecrete))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ExtractUserIdFromToken(c echo.Context) (uint, error) {
	tokenString := c.Request().Header.Get("Authorization")

	if tokenString == "" {
		return 0, jwt.ErrSignatureInvalid
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecrete), nil
	})

	if err != nil || !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, jwt.ErrSignatureInvalid
	}

	userID, ok := claim["user_id"].(float64)

	if !ok {
		return 0, jwt.ErrSignatureInvalid
	}

	return uint(userID), nil
}
