package test

import (
	"gateway/internal/middleware"
	"gateway/test"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

// Your Echo handler function
func MyHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}

func TestMyRoute(t *testing.T) {
	e, assert := test.EchoRouterProvider(t)

	// 2. Define your route
	middleware.HandlerMiddleware(e)

	// 3. Create a test server with your Echo instance
	ts := httptest.NewServer(e)
	defer ts.Close() // Important: Close the test server

	// 4. Make a request to the test server
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/auth/v1/registration", nil)
	assert.NoError(err)

	// 5. Use the test server's client to make the request
	client := ts.Client() // Use the test server's client for consistent behavior
	resp, err := client.Do(req)

	assert.NoError(err)
	defer resp.Body.Close() // Important: Close the response body

	// 6. Assertions on the response
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code mismatch")
	// Read the response body
	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)
	assert.NoError(err)

}
