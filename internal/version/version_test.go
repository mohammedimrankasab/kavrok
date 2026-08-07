// Package version provides build and runtime version information.
package version

import "testing"

func TestInfo(t *testing.T) {
	t.Parallel()

	info := Get()

	if info.Version == "" {
		t.Fatal("version should not be empty")
	}

	if info.Commit == "" {
		t.Fatal("commit should not be empty")
	}

	if info.BuildDate == "" {
		t.Fatal("build date should not be empty")
	}
}
