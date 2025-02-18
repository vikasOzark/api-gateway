package middleware

import (
	"gateway/internal/handler"
	"time"

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
	router.Use(middleware.RateLimiterWithConfig(RateLimiterConfig()))
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${remote_ip} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	router.Use(middleware.Recover())
	router.Use(handler.ExcludeRouteHandler)
	router.Use(AuthMiddleware)
	router.Use(handler.RouteHandler)
}

func FormatTimeRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}
