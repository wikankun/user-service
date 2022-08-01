package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwt_secret := os.Getenv("JWT_SECRET")
	return middleware.JWT([]byte(jwt_secret))
}
