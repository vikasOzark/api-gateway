package middleware

import (
	"gateway/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HandlerMiddleware sets up middleware for the Echo router.
//
// Parameters:
//   - router: An instance of *echo.Echo that represents the Echo router.
//
// The function performs the following steps:
// 1. Adds the RouteHandler middleware to the router to process and route
// incoming requests based on the HTTP method.
func HandlerMiddleware(router *echo.Echo) {
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(AuthMiddleware)
	router.Use(handler.RouteHandler)
}
