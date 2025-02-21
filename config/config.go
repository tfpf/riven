package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	FyneFont string `json:"FYNE_FONT,omitempty"`
}

// Read looks for the Riven configuration JSON file in the OS-specific user
// configuration directory and creates a configuration object.
func Read() *Config {
	cfg := &Config{}
	configDir, err := os.UserConfigDir()
	if err != nil {
		slog.Warn("Could not find user configuration directory", slog.Any("err", err))
		return cfg
	}
	configFile := filepath.Join(configDir, "riven", "config.json")
	configFileContents, err := os.ReadFile(configFile)
	if err != nil {
		slog.Warn("Could not read user configuration file", slog.Any("err", err), slog.String("file", configFile))
		return cfg
	}
	if err := json.Unmarshal(configFileContents, cfg); err != nil {
		slog.Warn("Could not parse user configuration file", slog.Any("err", err), slog.String("file", configFile))
		return cfg
	}
	slog.Info("Read user configuration file", slog.String("file", configFile), slog.Any("cfg", cfg))
	return cfg
}
