package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	FyneFont string `json:"FYNE_FONT"`
}

// Read looks for the Riven configuration JSON file in the OS-specific user
// configuration directory and loads it.
func Read() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		slog.Error("Could not find user configuration directory", slog.Any("err", err))
		return nil, err
	}
	configFile := filepath.Join(configDir, "riven", "config.json")
	configFileContents, err := os.ReadFile(configFile)
	cfg := &Config{}
	if err != nil {
		slog.Warn("Could not read user configuration file", slog.Any("err", err))
		return cfg, nil
	}
	if err := json.Unmarshal(configFileContents, cfg); err != nil {
		slog.Error("Could not parse user configuration file", slog.Any("err", err))
		return nil, err
	}
	if cfg.FyneFont != "" {
		// Tell Fyne to use this font.
		os.Setenv("FYNE_FONT", cfg.FyneFont)
	}
	slog.Info("Loaded user configuration", slog.Any("cfg", cfg))
	return cfg, nil
}
