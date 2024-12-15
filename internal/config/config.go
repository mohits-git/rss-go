package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".aggregatorconfig.json"

// Application configuration struct
type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// SetUser sets the current user in the config file
func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New(fmt.Sprintln("ERROR: @go-aggregator/internal/config: failed to get home directory\n", err))
	}
	return filepath.Join(homePath, configFileName), nil
}

// Reads .aggregatorconfig.json file and returns a Config struct
func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, errors.New(fmt.Sprintln("ERROR: @go-aggregator/internal/config: failed to open config file\n", err))
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, errors.New(fmt.Sprintln("ERROR: @go-aggregator/internal/config: failed to decode config\n", err))
	}

	return cfg, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(configFilePath, os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(fmt.Sprintln("ERROR: @go-aggregator/internal/config: failed to open config file\n", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(cfg); err != nil {
		return errors.New(fmt.Sprintln("ERROR: @go-aggregator/internal/config: failed to encode config\n", err))
	}
	return nil
}
