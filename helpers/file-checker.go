package helpers

import (
	"os"
)

// FileExists checks if a file exists in filesystem
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
