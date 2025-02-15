package test

import (
	"gateway/internal/middleware"
	"gateway/test"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestConcurrentRequests(t *testing.T) {
	e, assert := test.EchoRouterProvider(t)
	middleware.HandlerMiddleware(e)

	server := httptest.NewServer(e)
	defer server.Close()

	var wg sync.WaitGroup
	numRequests := 100
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
			assert.NoError(err)

			client := server.Client()
			resp, err := client.Do(req)

			assert.NoError(err)
			defer resp.Body.Close()

			assert.Equal(http.StatusOK, resp.StatusCode)

			body := new(strings.Builder)
			_, err = io.Copy(body, resp.Body)
			assert.NoError(err)
		}()
	}

	wg.Wait()
}
