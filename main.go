package main

import (
	"net/http"
	"os"
	"log"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
	"github.com/wikankun/user-service/database"
)

func main() {
	godotenv.Load()

	config :=
		database.Config{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			Port:     os.Getenv("DB_PORT"),
		}

	initDB(config)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			database.Migrate()
		}
	}

	startServer()
}

func initDB(config database.Config) {
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}

func initHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/register")
	e.POST("/verify")
	e.POST("/login")
}

func startServer() {
	port := os.Getenv("PORT")
	log.Printf("Starting HTTP Server on port %s", port)

	e := echo.New()
	initHandlers(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost},
	}))

	e.Logger.Fatal(e.Start(":"+port))
}