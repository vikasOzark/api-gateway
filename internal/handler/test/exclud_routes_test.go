package test

import (
	"fmt"
	"gateway/internal/middleware"
	"gateway/test"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExcludedRoutes(t *testing.T) {
	e, assert := test.EchoRouterProvider(t)
	middleware.HandlerMiddleware(e)

	testServer := httptest.NewServer(e)
	defer testServer.Close()

	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/auth/v1/registration", nil)
	assert.NoError(err)

	client := testServer.Client()
	resp, err := client.Do(req)

	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)

	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)
	assert.NoError(err)

	fmt.Printf("Response Body: %s", body.String())
}

func TestUnRegisteredRoutes(t *testing.T) {
	e, assert := test.EchoRouterProvider(t)

	middleware.HandlerMiddleware(e)

	testServer := httptest.NewServer(e)
	defer testServer.Close()

	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/test/not-found", nil)
	assert.NoError(err)

	client := testServer.Client()
	resp, err := client.Do(req)

	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusNotFound, resp.StatusCode)

	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)
	assert.NoError(err)
}
