package controller

import (
	"echo-api/service"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var otpStore = make(map[string]string)
var SigningKey = []byte("secret")

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	result, err := service.AuthenticateUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !result {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["email"] = email
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Generate encoded token and send it as response.
	t, err := token.SignedString(SigningKey)
	if err != nil {
		return err
	}
	// Membaca isi file HTML
	htmlContent, err := ioutil.ReadFile("./template/email.html")
	if err != nil {
		log.Println(htmlContent)
		log.Println("HTML Content Not Found")
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "admin123@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "From ")
	m.SetBody("text/html", string(htmlContent)) //parsing html

	// Inisialisasi objek Dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, "budimanfajar660@gmail.com", "bwtoyhwtdjwugbwq")

	// Kirim Email
	if err := d.DialAndSend(m); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func ForgotPassword(c echo.Context) error {
	email := c.FormValue("email")

	// Generate OTP
	rand.Seed(time.Now().UnixNano())
	otp := strconv.Itoa(rand.Intn(900000) + 100000)

	otpStore[email] = otp
	// Send OTP
	sendOTPEmail(email, otp)

	return c.JSON(http.StatusOK, "OTP sent to your email.")
}

func ResetPassword(c echo.Context) error {
	email := c.FormValue("email")
	otp := c.FormValue("otp")
	newPassword := c.FormValue("newPassword")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	hashed := string(hashedPassword)
	err := service.UpdatePassword(email, hashed)
	if err != nil {
		log.Println("Invalid Email or Password")
		return c.JSON(http.StatusBadRequest, err)
	}

	storedOTP, exists := otpStore[email]
	if !exists || storedOTP != otp {
		return c.JSON(http.StatusUnauthorized, "Invalid OTP.")
	}
	log.Printf("Password reset for email: %s. New password: %s\n", email, newPassword)

	// Hapus OTP
	delete(otpStore, email)

	return c.JSON(http.StatusOK, "Password reset successfully.")
}

func sendOTPEmail(email, otp string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset OTP")
	m.SetBody("text/plain", "Your OTP for password reset: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, "budimanfajar660@gmail.com", "bwtoyhwtdjwugbwq")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send OTP email: ", err)
	}
}
