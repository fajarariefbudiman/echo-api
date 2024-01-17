package controller

import (
	"echo-api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddOrders(c echo.Context) error {
	user_id := c.FormValue("userId")
	strusrid, err := strconv.Atoi(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	order_status := c.FormValue("orderStatus")
	payment_methood := c.FormValue("paymentMethod")
	result, err := service.AddOrders(strusrid, order_status, payment_methood)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, result)
}

func UpdateOrders(c echo.Context) error {
	orderid := c.Param("orderId")
	strid, err := strconv.Atoi(orderid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	result, err := service.UpdateOrders(strid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func GetOrderById(c echo.Context) error {
	id := c.Param("orderId")
	strid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Convert Error")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	result, err := service.GetOrderById(strid)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
