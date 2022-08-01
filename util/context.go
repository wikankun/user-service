package util

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TypeResponse struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
}

func MustGetIDFromContext(c echo.Context) int {
	contextKey := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	claims := contextKey.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}

func ErrorResponse(c echo.Context, code int, hint string) error {
	return c.JSON(code, TypeResponse{
		Success: false,
		Hint:    hint,
		Data:    nil,
	})
}

func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, TypeResponse{
		Success: true,
		Hint:    "",
		Data:    data,
	})
}
