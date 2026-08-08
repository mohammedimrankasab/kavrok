package kubernetes

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/version"
)

type fakeClient struct {
	serverVersion *version.Info
	err           error
}

func (f *fakeClient) ServerVersion(
	_ context.Context,
) (*version.Info, error) {
	return f.serverVersion, f.err
}

func TestClientServerVersion(t *testing.T) {
	t.Parallel()

	expected := &version.Info{
		Major:      "1",
		Minor:      "34",
		GitVersion: "v1.34.0",
	}

	client := &fakeClient{
		serverVersion: expected,
	}

	got, err := client.ServerVersion(context.Background())
	if err != nil {
		t.Fatalf("ServerVersion() failed: %v", err)
	}

	if got.GitVersion != expected.GitVersion {
		t.Fatalf(
			"expected version %q, got %q",
			expected.GitVersion,
			got.GitVersion,
		)
	}
}
