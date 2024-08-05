package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FileConfig struct {
		FilePath string `yaml:"file-name"`
	} `yaml:"file"`
}

// Loads and populates the configuration
// settings on the application
func LoadConfig(cfg *Config) error {
	f, err := os.Open("config.yaml")
	defer f.Close()
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		return err
	}
	return nil
}
