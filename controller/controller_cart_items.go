package controller

import (
	"echo-api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCartItems(c echo.Context) error {
	cart_id := c.Param("cartId")
	strcart, err := strconv.Atoi(cart_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Convert CartId Error",
		})
	}
	product_id := c.FormValue("productId")
	strprdc, err := strconv.Atoi(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Convert Product Id Error",
		})
	}
	quantity := c.FormValue("quantity")
	qtty, err := strconv.Atoi(quantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Convert Quantity Error",
		})
	}
	price := c.FormValue("price")
	if price == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Price is empty",
		})
	}

	strprice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Convert to float error")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Convert Price Error",
		})
	}
	result, err := service.AddCartItems(strcart, strprdc, qtty, float32(strprice))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, result)
}

func GetCartItems(c echo.Context) error {
	cart_id := c.Param("cartId")
	strcart, err := strconv.Atoi(cart_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Convert CartId Error",
		})
	}
	result, err := service.GetCartItems(strcart)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
