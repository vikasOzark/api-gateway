package main

import (
	"gateway/helpers"
	"os"
	"testing"
)

// TestMain sets up the test environment by loading the .env file
func TestMain(m *testing.M) {
	err := helpers.LoadEnvFile()
	if err != nil {
		panic("Failed to load .env file: " + err.Error())
	}

	// Run the tests
	code := m.Run()

	// Exit with the test status code
	os.Exit(code)
}

func TestRequiredCheck(t *testing.T) {
	// Call the RequiredChecks function and check for errors
	err := helpers.RequiredChecks()
	if err != nil {
		t.Errorf("RequiredChecks() returned an error: %v", err)
	}
}
