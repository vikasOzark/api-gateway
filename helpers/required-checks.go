package helpers

import (
	"fmt"
	"os"
	"strings"
)

// RequiredChecks performs necessary checks to ensure the environment and configuration
// files are properly set up. It loads the environment variables from a .env file,
// verifies the existence of the configuration file specified by the CONFIG_PATH
// environment variable, and ensures that the configuration file is in TOML format.
// If any of these checks fail, the function will return an error accordingly.
func RequiredChecks() error {
    if err := LoadEnvFile(); err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }

    configPath := os.Getenv("CONFIG_PATH")

    if configPath == "" {
        return fmt.Errorf("CONFIG_PATH environment variable not set")
    }

    if !strings.HasSuffix(configPath, ".toml") {
        return fmt.Errorf("config file must be in TOML format")
    }

    return nil
}