package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	log, err := New()

	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	if log == nil {
		t.Fatal("expected logger")
	}
}
