package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sversion "k8s.io/apimachinery/pkg/version"
)

func TestDoctorCommand(t *testing.T) {
	t.Parallel()

	cmd := newDoctorCommand(func() (kubernetes.Client, error) {
		return &fakeKubernetesClient{
			clusterName: "test-cluster",
			serverVersion: &k8sversion.Info{
				GitVersion: "v1.34.0",
			},
			nodes: &corev1.NodeList{Items: []corev1.Node{
				{ObjectMeta: metav1.ObjectMeta{Name: "test-node"},
					Status: corev1.NodeStatus{NodeInfo: corev1.NodeSystemInfo{KubeletVersion: "v1.34.0", OperatingSystem: "linux", Architecture: "arm64"}, Conditions: []corev1.NodeCondition{{Type: corev1.NodeReady, Status: corev1.ConditionTrue}}}}}},
			namespaceList: &corev1.NamespaceList{},
			pods:          &corev1.PodList{},
		}, nil
	})

	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)

	if err := cmd.Execute(); err != nil {
		t.Fatalf(
			"doctor command failed: %v\noutput:\n%s",
			err,
			output.String(),
		)
	}

	// Normalize multiple internal spaces into a single space for testing
	normalizedResult := strings.Join(strings.Fields(output.String()), " ")

	for _, expected := range []string{
		"Kavrok Doctor",
		"Cluster: test-cluster",
		"Version: v1.34.0",
		"Nodes: 1",
		"Namespaces: 0",
		"Pods: 0",
		"CLUSTER_HEALTHY",
		"No known health issues were detected.",
		"Result: PASSED",
	} {
		if !strings.Contains(normalizedResult, expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				output.String(), // Show the original output layout in the failure message
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

	// Normalize multiple internal spaces into a single space for testing
	normalizedResult := strings.Join(strings.Fields(output.String()), " ")

	for _, expected := range []string{"Kavrok Doctor", "Kubernetes unavailable", "connection refused", "Result: FAILED"} {
		if !strings.Contains(normalizedResult, expected) {
			t.Errorf(
				"expected output to contain %q, got %q",
				expected,
				output.String(), // Show the original output layout in the failure message
			)
		}
	}
}
