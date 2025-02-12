package requests

import (
	"fmt"
	"gateway/helpers"
	"net/http"
)

// RequestHandler sends a PUT request to the target location based on the provided URL and returns the response.
//
// Parameters:
//   - url: An instance of helpers.UrlExtractor that contains the context and service information.
//
// Returns:
//   - *http.Response: The response from the target location.
//   - error: An error if the request fails or the service is not found.
//
// The function performs the following steps:
//   1. Loads the configuration using the helpers.Config struct.
//   2. Retrieves the target location for the service from the configuration.
//   3. Extracts the payload and content type from the incoming request.
//   4. Creates a new HTTP client and a PUT request with the extracted payload and content type.
//   5. Sends the PUT request to the target location.
//   6. Returns the response from the target location or an error if the request fails.
func RequestHandler(url helpers.UrlExtractor) (*http.Response, error) {
    config := &helpers.Config{}

    err := config.LoadConfig()
    if err != nil {
        return nil, err
    }

    targetLocation := config.GetTarget(url.Service)
    if targetLocation == "" {
        return nil, fmt.Errorf("service not found")
    }

    payload := url.Context.Request().Body
    contentType := url.Context.Request().Header.Get(CONTENT_TYPE)

    client := &http.Client{}
    req, err := http.NewRequest(url.Method, url.GetTargetUrl(), payload)
    if err != nil {
        return nil, err
    }

    req.Header.Set(CONTENT_TYPE, contentType)
    response, err := client.Do(req)
    
    if err != nil {
        return nil, err
    }

    return response, nil
}
