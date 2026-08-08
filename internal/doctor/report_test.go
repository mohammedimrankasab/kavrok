package doctor

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteReportPassed(t *testing.T) {
	t.Parallel()

	result := Result{
		Checks: []Check{
			{
				Name:    "configuration",
				Status:  statusPass,
				Message: "configuration is valid",
			},
			{
				Name:    "kubernetes",
				Status:  statusPass,
				Message: "Kubernetes API is reachable",
			},
		},
	}

	var output bytes.Buffer

	if err := WriteReport(&output, result); err != nil {
		t.Fatalf("WriteReport() failed: %v", err)
	}

	for _, expected := range []string{
		"Kavrok Doctor",
		"✓",
		"configuration",
		"kubernetes",
		"Result: PASSED",
	} {
		if !strings.Contains(output.String(), expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				output.String(),
			)
		}
	}
}

func TestWriteReportFailed(t *testing.T) {
	t.Parallel()

	result := Result{
		Checks: []Check{
			{
				Name:    "configuration",
				Status:  statusPass,
				Message: "configuration is valid",
			},
			{
				Name:    "kubernetes",
				Status:  statusFail,
				Message: "Kubeconfig unavailable",
			},
		},
	}

	var output bytes.Buffer

	if err := WriteReport(&output, result); err != nil {
		t.Fatalf("WriteReport() failed: %v", err)
	}

	for _, expected := range []string{
		"Kavrok Doctor",
		"✓",
		"✗",
		"Kubeconfig unavailable",
		"Result: FAILED",
	} {
		if !strings.Contains(output.String(), expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				output.String(),
			)
		}
	}
}
