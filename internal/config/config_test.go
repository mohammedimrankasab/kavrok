package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Parallel()

	cfg := Load()

	if cfg.LogLevel != "info" {
		t.Fatalf("expected log level info, got %q", cfg.LogLevel)
	}

	if cfg.Output != "table" {
		t.Fatalf("expected output table, got %q", cfg.Output)
	}
}
