package controller

import (
	"echo-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeController(c echo.Context) error {
	result, err := service.HomeModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}
