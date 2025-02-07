package helpers

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Endpoints map[string]string `toml:"endpoints"`
	Exclude   map[string]string `toml:"exclude"`
}

func (cl *Config) LoadConfig() error {
	if _, err := toml.DecodeFile("config.toml", cl); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (cl *Config) GetTarget(target string) string {
	apiTarget := cl.Endpoints[target]
	return apiTarget
}
