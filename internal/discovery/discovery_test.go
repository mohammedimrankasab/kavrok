package discovery

import (
	"errors"
	"testing"

	"k8s.io/apimachinery/pkg/version"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)

func TestDiscover(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		clusterName: "test-cluster",
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
	if result.Name != "test-cluster" {
		t.Errorf(
			"expected cluster name %q, got %q",
			"test-cluster",
			result.Name,
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
		versionErr: errors.New("connection refused"),
	}

	_, err := Discover(client)
	if err == nil {
		t.Fatal("expected Discover() to fail")
	}

	if !errors.Is(err, client.versionErr) {
		t.Fatalf(
			"expected error to wrap %q, got %v",
			client.versionErr,
			err,
		)
	}
}
