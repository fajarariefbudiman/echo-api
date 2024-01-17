package controller

import (
	"database/sql"
	"echo-api/service"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	result, err := service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateProducts(c echo.Context) error {
	nama := c.FormValue("name")
	slug := c.FormValue("slug")
	price := c.FormValue("price")
	discount := c.FormValue("discount")
	strdisc, err := strconv.ParseFloat(discount, 32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	strprice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		log.Println("Convert to float error")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	stock := c.FormValue("stock")
	strstock, err := strconv.Atoi(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	category_id := c.FormValue("categoryId")
	strcategoryid, err := strconv.Atoi(category_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := service.CreateProducts(nama, slug, float32(strprice), float32(strdisc), strstock, strcategoryid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteProducts(c echo.Context) error {

	slug := c.FormValue("slug")
	result, err := service.DeleteProducts(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, result)
}

func UpdateProducts(c echo.Context) error {
	name := c.FormValue("name")
	slug := c.Param("slug")
	price := c.FormValue("price")
	strprice, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	stock := c.FormValue("stock")
	strstock, err := strconv.Atoi(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	categoryid := c.FormValue("categoryId")
	strcategory, err := strconv.Atoi(categoryid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := service.UpdateProducts(slug, name, strprice, strstock, strcategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetProductsBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := service.GetProductsSlug(slug)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
