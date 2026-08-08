// Package app contains the Kavrok application lifecycle and orchestration.
package app

import (
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/logger"
)

func TestNew(t *testing.T) {
	t.Parallel()

	log, err := logger.New()
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	app := New(log)

	if app == nil {
		t.Fatal("expected application to be created")
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()

	log, err := logger.New()
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	application := New(log)

	err = application.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

}
