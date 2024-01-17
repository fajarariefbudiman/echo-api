package controller

import (
	"echo-api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddOrderItems(c echo.Context) error {
	order_id := c.Param("orderId")
	strordrid, _ := strconv.Atoi(order_id)
	product_id := c.FormValue("productId")
	strprdctid, _ := strconv.Atoi(product_id)
	product_name := c.FormValue("productName")
	product_price := c.FormValue("productPrice")
	strprdtprice, _ := strconv.ParseFloat(product_price, 32)
	total_product := c.FormValue("totalProduct")
	strttlprdct, _ := strconv.Atoi(total_product)
	result, err := service.AddOrderItems(strordrid, strprdctid, product_name, float32(strprdtprice), strttlprdct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, result)
}
