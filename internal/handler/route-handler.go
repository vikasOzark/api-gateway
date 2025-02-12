package handler

import (
	"gateway/helpers"
	"gateway/internal/requests"

	"github.com/labstack/echo/v4"
)

// RouteHandler is a middleware function that processes incoming requests and routes them to the appropriate handler based on the HTTP method.
//
// Parameters:
//   - next: The next handler function in the middleware chain.
//
// Returns:
//   - echo.HandlerFunc: A function that processes the request and routes it to the appropriate handler.
//
// The function performs the following steps:
//  1. Extracts the URL and context information using helpers.UrlExtractor.
//  2. Processes the URL to determine the HTTP method.
//  3. Routes the request to the appropriate handler based on the HTTP method.
//     - If the method is "POST", it calls requests.HandlerPOST.
//  4. If the method is not recognized, it calls the next handler in the chain.
func RouteHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := helpers.UrlExtractor{
			Context: c,
		}
		// Process the URL to determine the HTTP method other useful information
		url.Process()
		return requests.RequestManager(url)
	}
}
