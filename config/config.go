package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// Config is the configuration for the application
type Config struct {
	// Name is the name of the application
	Name string

	// Version is the version of the application
	Version string

	// Author is the author of the application
	Author string

	// Email is the email of the author
	Email string

	// Description is the description of the application
	Description string

	// License is the license of the application
	License string

	// Path is the path to the application
	Path string

	// ConfigPath is the path to the configuration file
	ConfigPath string

	// ConfigName is the name of the configuration file
	ConfigName string

	// ConfigType is the type of the configuration file
	ConfigType string

	// LogPath is the path to the log file
	LogPath string
}

// New creates a new configuration
func New() *Config {
	return &Config{
		Name:        "routnd",
		Version:     "0.0.1",
		Author:      "DeeStarks",
		Email:       "danielonoja246@gmail.com",
		Description: "A CLI tool for executing commands/processes in the background",
		License:     "MIT",
		Path:        getPath(),
		ConfigPath:  getConfigPath(),
		ConfigName:  "config",
		ConfigType:  "yaml",
		LogPath:     getLogPath(),
	}
}

// getPath returns the path to the application
func getPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// getConfigPath returns the path to the configuration file
func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".routnd", "config")
}

// getLogPath returns the path to the log file
func getLogPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".routnd", "logs")
}
