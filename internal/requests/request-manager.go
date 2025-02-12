package requests

import (
	"gateway/helpers"
	"io"
	"net/http"
)

var (
	CONTENT_TYPE = "Content-Type"
)

// RequestManager sends a request to a target location based on the provided URL and returns the response to the client.
//
// Parameters:
//   - url: An instance of helpers.UrlExtractor that contains the context and service information.
//
// Returns:
//   - error: An error if the request fails or the service is not found.
//
// The function performs the following steps:
//  1. Extracts the payload and content type from the incoming request.
//  2. Loads the configuration using the helpers.Config struct.
//  3. Retrieves the target location for the service from the configuration.
//  4. If the target location is not found, returns a 404 Not Found error.
//  5. Sends a POST request to the target location with the extracted payload and content type.
//  6. If the POST request fails, returns a 500 Internal Server Error with the error message.
//  7. Sets the response content type to match the target's response content type.
//  8. Reads the response body from the target location.
//  9. Returns the response body to the client with the appropriate content type.
func RequestManager(url helpers.UrlExtractor) error {

	cl := &helpers.Config{}
	cl.LoadConfig()
	targetLocation := cl.GetTarget(url.Service)

	if targetLocation == "" {
		return url.Context.JSON(http.StatusNotFound, map[string]string{"ok": "false", "error 3": "Service not found"})
	}

	// Send POST request to the target location
	response, err := RequestHandler(url)
	if err != nil {
		return url.Context.JSON(http.StatusInternalServerError, map[string]string{"ok": "false", "error 1": err.Error()})
	}
	defer response.Body.Close()

	// Get the response content type
	responseContentType := response.Header.Get(CONTENT_TYPE)

	// Set the response content type
	url.Context.Response().Header().Set(CONTENT_TYPE, responseContentType)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return url.Context.JSON(response.StatusCode, map[string]string{"ok": "false", "error 2": err.Error()})
	}

	// Return the response body to the client
	return helpers.ContentTypeResponse(url.Context, responseContentType, body, response)
}