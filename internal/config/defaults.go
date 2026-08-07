// Package config provides application configuration handling.
package config

// Defaults returns the default configuration for Kavrok.
func Defaults() Config {
	return Config{
		LogLevel: "info",
		Output:   "table",
	}
}
