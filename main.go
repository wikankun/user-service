package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wikankun/user-service/config"
	"github.com/wikankun/user-service/database"
	"github.com/wikankun/user-service/migration"
	"github.com/wikankun/user-service/router"
	"github.com/wikankun/user-service/util"
)

func main() {
	godotenv.Load()

	config.InitConfig()
	database.InitDB()
	util.InitUtil()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			migration.MigrateDB()
		}
	}

	startServer()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func startServer() {
	port := config.Config.App.Port
	log.Printf("Starting HTTP Server on port %s", port)

	e := echo.New()

	router.InitHandlers(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
	}))
	e.Use(middleware.Logger())

	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	e.Logger.Fatal(e.Start(":" + port))
}
