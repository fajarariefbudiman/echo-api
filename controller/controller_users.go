package controller

import (
	"database/sql"
	"echo-api/service"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func GetUsers(c echo.Context) error {
	result, err := service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateUsers(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	handphone := c.FormValue("handphone")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashed := string(hashedPassword)
	result, err := service.CreateUsers(name, email, hashed, handphone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	htmlcontent, _ := ioutil.ReadFile("../template/email.html")
	m := gomail.NewMessage()
	m.SetHeader("From", "admin123@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "From ")
	m.SetBody("text/html", string(htmlcontent)) //parsing html

	// Inisialisasi objek Dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, "budimanfajar660@gmail.com", "bwtoyhwtdjwugbwq")

	// Kirim Email
	if err := d.DialAndSend(m); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteUsers(c echo.Context) error {

	id := c.FormValue("id")
	result, err := service.DeleteProducts(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, result)
}

func UpdateUsers(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	handphone := c.FormValue("handphone")
	id := c.Param("id")
	strid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result, err := service.UpdateUsers(strid, name, email, password, handphone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetUsersId(c echo.Context) error {
	id := c.Param("id")
	strid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result, err := service.GetUsersId(strid)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAddressesUsers(c echo.Context) error {
	user_name := c.Param("userName")
	result, err := service.GetAddressesUsers(user_name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, result)
}
