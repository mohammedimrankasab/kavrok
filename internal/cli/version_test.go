package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	t.Parallel()

	cmd := newVersionCommand()

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("version command failed: %v", err)
	}

	result := output.String()

	for _, expected := range []string{
		"Kavrok",
		"Commit:",
		"Build Date:",
		"Go Version:",
		"Platform:",
		"Tree State:",
	} {
		if !strings.Contains(result, expected) {
			t.Errorf("expected output to contain %q, got %q", expected, result)
		}
	}
}

func TestVersionCommandShort(t *testing.T) {
	t.Parallel()

	cmd := newVersionCommand()

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{"--short"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("version command failed: %v", err)
	}

	result := strings.TrimSpace(output.String())

	if result != "dev" {
		t.Fatalf("expected short version %q, got %q", "dev", result)
	}
}
