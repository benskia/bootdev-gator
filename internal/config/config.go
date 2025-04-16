package config

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/benskia/Gator/internal/errors"
)

const configFileName string = ".gatorconfig.json"

// Config struct    Models a json config file.
type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read function    Returns a Config struct that models a json config file.
func Read() (*Config, error) {
	errTagger := errors.NewErrTagger("Read " + configFileName)

	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, errTagger(err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errTagger(err)
	}

	newCfg := Config{}
	if err := json.Unmarshal(data, &newCfg); err != nil {
		return nil, errTagger(err)
	}

	return &newCfg, nil
}

// SetUser method    Sets cfg.CurrentUserName to username, and writes the
// Config to .gatorconfig.json at user's home directory.
func (cfg *Config) SetUser(username string) error {
	errTagger := errors.NewErrTagger("SetUser " + username)

	log.Printf("setting user %s\n", username)
	cfg.CurrentUserName = username
	if err := writeConfig(*cfg); err != nil {
		return errTagger(err)
	}

	return nil
}

// getConfigFilePath function    Returns a filepath assuming a config file at
// user's home directory.
func getConfigFilePath() (string, error) {
	errTagger := errors.NewErrTagger("getConfigFilePath")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errTagger(err)
	}

	return path.Join(homeDir, configFileName), nil
}

// writeConfig function    Writes cfg to file as json.
func writeConfig(cfg Config) error {
	errTagger := errors.NewErrTagger("writeConfig")

	data, err := json.Marshal(cfg)
	if err != nil {
		return errTagger(err)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return errTagger(err)
	}

	if err := os.WriteFile(configFilePath, data, 0644); err != nil {
		return errTagger(err)
	}

	return nil
}
