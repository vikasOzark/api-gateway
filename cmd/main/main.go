package main

import (
	"gateway/helpers"
	"gateway/internal/middleware"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func init() {
	helpers.RequiredChecks()
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
