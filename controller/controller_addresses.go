package controller

import (
	"echo-api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddAddressesUsers(c echo.Context) error {
	user_id := c.FormValue("userId")
	strid, err := strconv.Atoi(user_id)
	if err != nil {
		log.Println("Convert Error")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	user_name := c.FormValue("userName")
	street := c.FormValue("streetAddresses")
	city := c.FormValue("city")
	province := c.FormValue("province")
	country := c.FormValue("country")
	phone_number := c.FormValue("phoneNumber")
	result, err := service.CreateAddresses(strid, user_name, street, city, province, country, phone_number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, result)
}
