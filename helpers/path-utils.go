package helpers

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// GetProjectRoot returns the root path of the project by walking up the directory tree
// until it finds the specified marker file or directory (e.g., .git, go.mod).
func GetProjectRoot(marker string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd, marker)); err == nil {
			return cwd, nil
		}
		parent := filepath.Dir(cwd)
		if parent == cwd {
			break
		}
		cwd = parent
	}
	return "", os.ErrNotExist
}

// LoadEnvFile loads the .env file from the project root.
func LoadEnvFile() error {
	rootPath, err := GetProjectRoot("go.mod")
	if err != nil {
		return err
	}

	envPath := filepath.Join(rootPath, ".env")
	return godotenv.Load(envPath)
}
