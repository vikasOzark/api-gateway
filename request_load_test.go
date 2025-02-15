package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestLoadRequest(t *testing.T) {
	// Configuration
	url := "http://localhost:3002/user/v1/users"
	numRequests := 100 // Number of concurrent requests
	numWorkers := 10   // Number of goroutines

	var wg sync.WaitGroup
	requests := make(chan struct{}, numRequests)
	completed := make(chan struct{}, numRequests)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go worker(url, requests, completed, &wg)
	}

	// Enqueue requests
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		requests <- struct{}{}
	}
	close(requests)

	// Monitor progress
	go func() {
		for i := 0; i < numRequests; i++ {
			<-completed
			fmt.Printf("\rCompleted %d/%d requests", i+1, numRequests)
		}
		fmt.Println("\nLoad test completed")
	}()

	// Wait for all requests to complete
	wg.Wait()
}

func worker(url string, requests <-chan struct{}, completed chan<- struct{}, wg *sync.WaitGroup) {
	for range requests {
		sendRequest(url)
		wg.Done()
		completed <- struct{}{}
	}
}

func sendRequest(url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}

	body := new(strings.Builder)
	_, err = io.Copy(body, resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
}
