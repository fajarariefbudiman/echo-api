package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(Next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")
		if apiKey != "RAHASIA" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "UNAUTHORIZED"})
		}
		return Next(c)
	}
}
