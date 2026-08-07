package config

func Defaults() Config {
	return Config{
		LogLevel: "info",
		Output:   "table",
	}
}
