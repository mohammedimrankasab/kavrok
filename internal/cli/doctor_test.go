package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	k8sversion "k8s.io/apimachinery/pkg/version"
)

func TestDoctorCommand(t *testing.T) {
	t.Parallel()

	cmd := newDoctorCommand(func() (kubernetes.Client, error) {
		return &fakeKubernetesClient{
			serverVersion: &k8sversion.Info{
				GitVersion: "v1.34.0",
			},
		}, nil
	})

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("doctor command failed: %v", err)
	}

	result := output.String()

	for _, expected := range []string{
		"Kavrok Doctor",
		"configuration",
		"configuration is valid",
		"kubernetes",
		"Kubernetes API is reachable",
		"Result: PASSED",
	} {
		if !strings.Contains(result, expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				result,
			)
		}
	}
}

func TestDoctorCommandKubernetesFailure(t *testing.T) {
	t.Parallel()

	cmd := newDoctorCommand(func() (kubernetes.Client, error) {
		return nil, errors.New("connection refused")
	})

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected doctor command to fail")
	}

	result := output.String()

	for _, expected := range []string{
		"Kavrok Doctor",
		"configuration is valid",
		"kubeconfig unavailable",
		"connection refused",
		"Result: FAILED",
	} {
		if !strings.Contains(result, expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				result,
			)
		}
	}
}
