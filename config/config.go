package config

import (
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	FyneFont string `json:"FYNE_FONT"`
}

// NewConfig looks for the Riven configuration JSON file in the OS-specific
// user configuration directory and loads it.
func NewConfig() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		slog.Error("Failed to find user configuration directory", slog.Any("err", err))
		return nil, err
	}
	configFile := filepath.Join(configDir, "riven", "config.json")
	_, err = os.Open(configFile)
	if err != nil {
		slog.Warn("Failed to load user configuration file", slog.Any("err", err))
		return nil, err
	}
	return nil, nil
}
