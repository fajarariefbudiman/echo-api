package controller

import (
	"database/sql"
	"echo-api/service"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetAllCategories(c echo.Context) error {
	result, err := service.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/categories",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, result)
}

func CreateCategories(c echo.Context) error {
	nama := c.FormValue("name")
	slug := c.FormValue("slug")
	result, err := service.CreateCategories(nama, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteCategories(c echo.Context) error {
	id := c.FormValue("id")
	strid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result, err := service.DeleteCategories(strid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, result)
}

func UpdateCategories(c echo.Context) error {
	id := c.Param("id")
	strid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	slug := c.FormValue("slug")
	result, err := service.UpdateCategories(strid, name, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetCategoriesId(c echo.Context) error {
	slug := c.Param("slug")
	result, err := service.GetCategoriesSlug(slug)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
