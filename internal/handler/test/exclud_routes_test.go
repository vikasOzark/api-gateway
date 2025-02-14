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

	ts := httptest.NewServer(e)
	defer ts.Close() 

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/auth/v1/registration", nil)
	assert.NoError(err)

	client := ts.Client() 
	resp, err := client.Do(req)

	assert.NoError(err)
	defer resp.Body.Close() 

	assert.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	
	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)
	fmt.Println(body)
	assert.NoError(err)
}

func TestUnRegisteredRoutes(t *testing.T) {
	e, assert := test.EchoRouterProvider(t)

	middleware.HandlerMiddleware(e)

	ts := httptest.NewServer(e)
	defer ts.Close() 

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/test/not-found", nil)
	assert.NoError(err)

	client := ts.Client() 
	resp, err := client.Do(req)

	assert.NoError(err)
	defer resp.Body.Close() 

	assert.Equal(http.StatusNotFound, resp.StatusCode)
	
	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)

	if !assert.Equal(http.StatusNotFound, resp.StatusCode) {
		fmt.Println(resp.Body)
	}
	
	assert.NoError(err)
}