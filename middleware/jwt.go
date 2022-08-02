package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wikankun/user-service/config"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwt_secret := config.Config.App.JWTSecret
	return middleware.JWT([]byte(jwt_secret))
}
