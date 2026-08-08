package doctor

import (
	"errors"
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"k8s.io/apimachinery/pkg/version"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	err           error
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	return f.serverVersion, f.err
}

func TestResultPassed(t *testing.T) {
	t.Parallel()

	result := Result{
		Checks: []Check{
			{
				Name:   "configuration",
				Status: statusPass,
			},
			{
				Name:   "kubernetes",
				Status: statusPass,
			},
		},
	}

	if !result.Passed() {
		t.Fatal("expected result to pass")
	}
}

func TestResultFailed(t *testing.T) {
	t.Parallel()

	result := Result{
		Checks: []Check{
			{
				Name:   "configuration",
				Status: statusPass,
			},
			{
				Name:   "kubernetes",
				Status: statusFail,
			},
		},
	}

	if result.Passed() {
		t.Fatal("expected result to fail")
	}
}

func TestRunnerRunKubernetesAvailable(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		serverVersion: &version.Info{
			GitVersion: "v1.34.0",
		},
	}

	result := New(func() (kubernetes.Client, error) {
		return client, nil
	}).Run()

	if !result.Passed() {
		t.Fatal("expected diagnostics to pass")
	}

	if len(result.Checks) != 2 {
		t.Fatalf("expected 2 checks, got %d", len(result.Checks))
	}

	kubernetesCheck := result.Checks[1]

	if kubernetesCheck.Status != statusPass {
		t.Fatalf(
			"expected kubernetes check to pass, got %q",
			kubernetesCheck.Status,
		)
	}
}

func TestRunnerRunKubernetesUnavailable(t *testing.T) {
	t.Parallel()

	result := New(func() (kubernetes.Client, error) {
		return nil, errors.New("kubeconfig not found")
	}).Run()

	if result.Passed() {
		t.Fatal("expected diagnostics to fail")
	}

	kubernetesCheck := result.Checks[1]

	if kubernetesCheck.Status != statusFail {
		t.Fatalf(
			"expected kubernetes check to fail, got %q",
			kubernetesCheck.Status,
		)
	}
}
