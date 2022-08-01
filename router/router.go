package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wikankun/user-service/controller"
	mw "github.com/wikankun/user-service/middleware"
)

func InitHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/register", controller.UserRegister)
	e.POST("/verify", controller.UserVerify)
	e.POST("/login", controller.UserGetToken)
	e.GET("/user", controller.UserInfo, mw.JWTMiddleware())
	e.PUT("/user", controller.UserUpdate, mw.JWTMiddleware())
}
