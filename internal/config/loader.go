// Package config provides application configuration handling.
package config

// Load loads the application configuration.
func Load() Config {
	return Defaults()
}
