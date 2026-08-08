package discovery

import (
	"errors"
	"testing"

	"k8s.io/apimachinery/pkg/version"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	err           error
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.serverVersion, nil
}

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)

func TestDiscover(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		serverVersion: &version.Info{
			GitVersion: "v1.34.0",
			Platform:   "linux/arm64",
			GitCommit:  "abc123",
		},
	}

	result, err := Discover(client)
	if err != nil {
		t.Fatalf("Discover() failed: %v", err)
	}

	if result.Version != "v1.34.0" {
		t.Errorf(
			"expected version %q, got %q",
			"v1.34.0",
			result.Version,
		)
	}

	if result.Platform != "linux/arm64" {
		t.Errorf(
			"expected platform %q, got %q",
			"linux/arm64",
			result.Platform,
		)
	}

	if result.Commit != "abc123" {
		t.Errorf(
			"expected commit %q, got %q",
			"abc123",
			result.Commit,
		)
	}
}

func TestDiscoverServerVersionFailure(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		err: errors.New("connection refused"),
	}

	_, err := Discover(client)
	if err == nil {
		t.Fatal("expected Discover() to fail")
	}

	if !errors.Is(err, client.err) {
		t.Fatalf(
			"expected error to wrap %q, got %v",
			client.err,
			err,
		)
	}
}
