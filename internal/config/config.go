package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/benskia/Gator/internal/gatorErrs"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read() function    Returns a Config struct that models a json config file.
func Read() (*Config, error) {
	errWrap := gatorerrs.NewErrWrapper("Read")

	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, errWrap("failed to get config filepath", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errWrap("failed to read config filed", err)
	}

	newCfg := Config{}
	if err := json.Unmarshal(data, &newCfg); err != nil {
		return nil, errWrap("failed to unmarshal data", err)
	}

	return &newCfg, nil
}

// SetUser() method    Sets cfg.CurrentUserName to username, and writes the
// Config to .gatorconfig.json at user's home directory.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	if err := writeConfig(*cfg); err != nil {
		return fmt.Errorf("SetUser: %w", err)
	}
	return nil
}

// getConfigFilePath() function    Returns a filepath assuming a config file at
// user's home directory.
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getConfigFilePath: %w", err)
	}
	return path.Join(homeDir, configFileName), nil
}

// writeConfig() function    Writes cfg to file as json.
func writeConfig(cfg Config) error {
	errWrap := gatorerrs.NewErrWrapper("writeConfig")

	data, err := json.Marshal(cfg)
	if err != nil {
		return errWrap("failed to marshal data", err)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return errWrap("failed to get config filepath", err)
	}

	if err := os.WriteFile(configFilePath, data, 0644); err != nil {
		return errWrap("failed to write config file", err)
	}

	return nil
}
