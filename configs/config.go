package configs

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

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

	todosFilePath := ""

	switch runtime.GOOS {
	case "windows":
		todosFilePath = filepath.Join(os.Getenv("APPDATA"), "r2todo", "todos.csv")

	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		todosFilePath = filepath.Join(homeDir, ".config", "r2todo", "todos.csv")

	case "linux":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		todosFilePath = filepath.Join(homeDir, ".config", "r2todo", "todos.csv")
	default:
		return errors.New("Unsupported platform")

	}
	dir := filepath.Dir(todosFilePath)
	err = os.MkdirAll(dir, 0755) // 0755 is a common permission setting
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.OpenFile(todosFilePath, os.O_RDWR, 0644)
	defer file.Close()
	if os.IsNotExist(err) {
		f, err := os.Create(todosFilePath)
		defer f.Close()
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	cfg.FileConfig.FilePath = todosFilePath
	return nil
}
