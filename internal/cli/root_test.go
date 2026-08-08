package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewRootCommand(t *testing.T) {
	t.Parallel()

	cmd := NewRootCommand(newTestDependencies())

	if cmd == nil {
		t.Fatal("expected root command")
	}

	if cmd.Use != "kavrok" {
		t.Fatalf("expected command name kavrok, got %q", cmd.Use)
	}
}

func TestRootCommandHelp(t *testing.T) {
	t.Parallel()

	cmd := NewRootCommand(newTestDependencies())

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{"--help"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("help command failed: %v", err)
	}

	result := output.String()

	for _, expected := range []string{
		"Kubernetes engineering intelligence CLI",
		"version",
		"doctor",
	} {
		if !strings.Contains(result, expected) {
			t.Errorf(
				"expected help output to contain %q, got %q",
				expected,
				result,
			)
		}
	}
}

func TestRootCommandUnknownCommand(t *testing.T) {
	t.Parallel()

	cmd := NewRootCommand(newTestDependencies())
	cmd.SetArgs([]string{"unknown"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected unknown command to return an error")
	}
}
