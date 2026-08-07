package app

import "testing"

func TestNew(t *testing.T) {
	t.Parallel()

	app := New()

	if app == nil {
		t.Fatal("expected application to be created")
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()

	app := New()

	if err := app.Execute(); err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}
}
