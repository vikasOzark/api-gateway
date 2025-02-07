package middleware

import "github.com/labstack/echo/v4"

type AuthRequest struct {
	RequestEndpoint string `json:"request_endpoint"`
	RequestScope    string `json:"request_scope"`
	ServiceName     string `json:"registered_service_name"`
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Perform authentication here
		return next(c)
	}

}
