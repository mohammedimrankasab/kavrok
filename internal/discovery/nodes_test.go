package discovery

import (
	"context"
	"errors"
	"strings"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDiscoverNodes(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		nodes: &corev1.NodeList{
			Items: []corev1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-1",
						Labels: map[string]string{
							"node-role.kubernetes.io/control-plane": "",
						},
					},
					Status: corev1.NodeStatus{
						NodeInfo: corev1.NodeSystemInfo{
							KubeletVersion:  "v1.34.0",
							OperatingSystem: "linux",
							Architecture:    "arm64",
						},
						Conditions: []corev1.NodeCondition{
							{
								Type:   corev1.NodeReady,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
			},
		},
	}

	result, err := DiscoverNodes(context.Background(), client)
	if err != nil {
		t.Fatalf("DiscoverNodes() failed: %v", err)
	}

	if len(result.Nodes) != 1 {
		t.Fatalf(
			"expected 1 node, got %d",
			len(result.Nodes),
		)
	}

	node := result.Nodes[0]

	if node.Name != "node-1" {
		t.Errorf(
			"expected node name %q, got %q",
			"node-1",
			node.Name,
		)
	}

	if !node.Ready {
		t.Error("expected node to be ready")
	}

	if node.KubernetesVersion != "v1.34.0" {
		t.Errorf(
			"expected Kubernetes version %q, got %q",
			"v1.34.0",
			node.KubernetesVersion,
		)
	}

	if node.OS != "linux" {
		t.Errorf(
			"expected OS %q, got %q",
			"linux",
			node.OS,
		)
	}

	if node.Architecture != "arm64" {
		t.Errorf(
			"expected architecture %q, got %q",
			"arm64",
			node.Architecture,
		)
	}

	if len(node.Roles) != 1 || node.Roles[0] != "control-plane" {
		t.Errorf(
			"expected control-plane role, got %v",
			node.Roles,
		)
	}
}

func TestDiscoverNodesFailure(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		nodesErr: errors.New("connection refused"),
	}

	_, err := DiscoverNodes(context.Background(), client)
	if err == nil {
		t.Fatal("expected DiscoverNodes() to fail")
	}

	if !strings.Contains(err.Error(), "list Kubernetes nodes") {
		t.Fatalf(
			"expected wrapped node error, got %v",
			err,
		)
	}
}

func TestDiscoverNodesCapturesAllocatableResources(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		nodes: &corev1.NodeList{
			Items: []corev1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node",
					},
					Status: corev1.NodeStatus{
						NodeInfo: corev1.NodeSystemInfo{
							KubeletVersion:  "v1.36.1",
							OperatingSystem: "linux",
							Architecture:    "arm64",
						},
						Allocatable: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("8"),
							corev1.ResourceMemory: resource.MustParse("7Gi"),
						},
					},
				},
			},
		},
	}

	result, err := DiscoverNodes(context.Background(), client)
	if err != nil {
		t.Fatalf("DiscoverNodes() failed: %v", err)
	}

	if len(result.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(result.Nodes))
	}

	node := result.Nodes[0]

	if node.AllocatableCPU != "8" {
		t.Errorf(
			"expected allocatable CPU %q, got %q",
			"8",
			node.AllocatableCPU,
		)
	}

	if node.AllocatableMemory != "7Gi" {
		t.Errorf(
			"expected allocatable memory %q, got %q",
			"7Gi",
			node.AllocatableMemory,
		)
	}
}
