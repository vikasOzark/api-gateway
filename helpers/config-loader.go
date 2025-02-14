package helpers

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config represents the configuration structure for the API gateway.
// It contains two fields:
// - Endpoints: a map of endpoint names to their corresponding URLs, tagged for TOML parsing.
// - Exclude: a map of endpoint names to be excluded, tagged for TOML parsing.
type Config struct {
	Endpoints map[string]string `toml:"endpoints"`
	Exclude   map[string]string `toml:"exclude"`
}

func (cl *Config) LoadConfig() error {
	config_path := os.Getenv("CONFIG_PATH")

	if _, err := toml.DecodeFile(config_path, cl); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (cl *Config) GetTarget(target string) string {
	apiTarget := cl.Endpoints[target]
	return apiTarget
}

// IsExcluded checks if the given target is in the Exclude map of the Config.
// It returns the value associated with the target and a boolean indicating
// whether the target was found in the Exclude map.
//
// Parameters:
//
//	target - the key to look for in the Exclude map
//
// Returns:
//
//	value - the value associated with the target if found
//	ok - true if the target is found in the Exclude map, false otherwise
func (cl *Config) IsExcluded(target string) (string, bool) {
	value, ok := cl.Exclude[target]
	return value, ok
}
