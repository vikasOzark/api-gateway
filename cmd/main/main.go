package main

import (
	"gateway/helpers"
	"gateway/internal/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	error_ := godotenv.Load()
	if error_ != nil {
		log.Fatal("Unable to load .env file : ", error_)
	}

	config_path := os.Getenv("CONFIG_PATH")
	isExists := helpers.FileExists(config_path)

	if !isExists {
		panic("Config file not found")
	}

	isToml := strings.HasSuffix(config_path, ".toml")
	if !isToml {
		panic("Config file must be in TOML format")
	}

}

func main() {
	helpers.Logger().Info("Starting the API gateway service")
	
	debugEnv := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugEnv)
	if err != nil {
		log.Printf("Invalid value for DEBUG: %s, defaulting to false", debugEnv)
		debug = false
	}

	router := echo.New()
	router.Debug = debug

	// This helper function is used to add the middlewares to the router.
	middleware.HandlerMiddleware(router)

	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<body style="background-color:black; height:100vh;">
	<h1 style="color:#AEEA94;">Welcome to the API gateway service.</h1></body>`)
	})

	severPort := os.Getenv("PORT")
	router.Logger.Fatal(router.Start(":" + severPort))
}
