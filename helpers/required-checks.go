package helpers

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// RequiredChecks performs necessary checks to ensure the environment and configuration
// files are properly set up. It loads the environment variables from a .env file,
// verifies the existence of the configuration file specified by the CONFIG_PATH
// environment variable, and ensures that the configuration file is in TOML format.
// If any of these checks fail, the function will log an error or panic accordingly.
func RequiredChecks() {
	envValue := os.Getenv("ISDELVE_Enabled")
	isEnabled, _ := strconv.ParseBool(envValue)
	
	projectRoot := ".env"
	if isEnabled {
		projectRoot = filepath.Join("..", "..", ".env")
	}

	if err := godotenv.Load(projectRoot); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config_path := os.Getenv("CONFIG_PATH")
	if isEnabled {
		config_path = filepath.Join("..", "..", "config.toml")
	}

	isExists := FileExists(config_path)
	if !isExists {
		panic("Config file not found")
	}

	isToml := strings.HasSuffix(config_path, ".toml")
	if !isToml {
		panic("Config file must be in TOML format")
	}
}
