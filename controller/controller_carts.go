package controller

import (
	"database/sql"
	"echo-api/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func AddCarts(c echo.Context) error {
	cart_id := c.FormValue("userId")
	strid, err := strconv.Atoi(cart_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	result, err := service.AddCarts(strid)
	cookie := new(http.Cookie)
	cookie.Name = "name"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(http.StatusCreated, result)
}

func GetCartByUserId(c echo.Context) error {
	id := c.Param("userId")
	strid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Convert Error")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	result, err := service.GetCartByUserId(strid)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"Message": "No data found for the specified user ID",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
