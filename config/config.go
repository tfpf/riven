package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

// Config stores Riven configuration parameters.
type Config struct {
	configFile string
	FyneFont   string `json:"fyneFont,omitempty"`
}

// locate sets the location where the Riven configuration JSON file should be
// on this OS if it is not already set.
func (cfg *Config) locate() error {
	if cfg.configFile != "" {
		return nil
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		slog.Error("Could not find user configuration directory", slog.Any("err", err))
		return err
	}
	cfg.configFile = filepath.Join(configDir, "riven", "config.json")
	return nil
}

// Read reads Riven configuration parameters from the Riven configuration JSON
// file in the OS-specific user configuration directory.
func (cfg *Config) Read() error {
	if err := cfg.locate(); err != nil {
		return err
	}
	configFileContents, err := os.ReadFile(cfg.configFile)
	if err != nil {
		slog.Error("Could not read Riven configuration file", slog.Any("err", err), slog.String("file", cfg.configFile))
		return err
	}
	if err := json.Unmarshal(configFileContents, cfg); err != nil {
		slog.Error("Could not decode Riven configuration file", slog.Any("err", err), slog.String("file", cfg.configFile))
		return err
	}
	slog.Info("Read Riven configuration file", slog.String("file", cfg.configFile), slog.Any("cfg", cfg))
	return nil
}

// Write writes Riven configuration parameters to the Riven configuration JSON
// file in the OS-specific user configuration directory.
func (cfg *Config) Write() error {
	if err := cfg.locate(); err != nil {
		return err
	}
	configFileContents, err := json.Marshal(cfg)
	if err != nil {
		slog.Error("Could not encode Riven configuration", slog.Any("err", err), slog.Any("cfg", cfg))
		return err
	}
	if err := os.WriteFile(cfg.configFile, configFileContents, 0644); err != nil {
		slog.Error("Could not write Riven configuration file", slog.Any("err", err), slog.String("file", cfg.configFile))
		return err
	}
	slog.Info("Wrote Riven configuration file", slog.String("file", cfg.configFile), slog.Any("cfg", cfg))
	return nil
}
