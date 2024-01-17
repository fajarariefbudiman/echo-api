package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var AuthLogin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
