package discovery

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
)

func TestDiscoverSnapshot(t *testing.T) {
	t.Parallel()

	client := &fakeKubernetesClient{
		serverVersion: &version.Info{
			GitVersion: "v1.34.0",
			Platform:   "linux/arm64",
			GitCommit:  "abc123",
		},
		nodes: &corev1.NodeList{
			Items: []corev1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-1",
					},
					Status: corev1.NodeStatus{
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
		namespaces: &corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "kube-system",
					},
				},
			},
		},
		pods: &corev1.PodList{
			Items: []corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "api",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
						Conditions: []corev1.PodCondition{
							{
								Type:   corev1.PodReady,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "worker",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodPending,
					},
				},
			},
		},
	}

	result, err := DiscoverSnapshot(
		context.Background(),
		client,
	)
	if err != nil {
		t.Fatalf("DiscoverSnapshot() failed: %v", err)
	}

	if result.Cluster.Version != "v1.34.0" {
		t.Errorf(
			"expected cluster version %q, got %q",
			"v1.34.0",
			result.Cluster.Version,
		)
	}

	if len(result.Nodes.Nodes) != 1 {
		t.Errorf(
			"expected 1 node, got %d",
			len(result.Nodes.Nodes),
		)
	}

	if len(result.Namespaces) != 2 {
		t.Errorf(
			"expected 2 namespaces, got %d",
			len(result.Namespaces),
		)
	}

	if result.Workloads.PodCount != 2 {
		t.Errorf(
			"expected 2 pods, got %d",
			result.Workloads.PodCount,
		)
	}

	if result.Workloads.RunningPods != 1 {
		t.Errorf(
			"expected 1 running pod, got %d",
			result.Workloads.RunningPods,
		)
	}

	if result.Workloads.PendingPods != 1 {
		t.Errorf(
			"expected 1 pending pod, got %d",
			result.Workloads.PendingPods,
		)
	}

	if result.Workloads.ReadyPods != 1 {
		t.Errorf(
			"expected 1 ready pod, got %d",
			result.Workloads.ReadyPods,
		)
	}

	if !result.Health.Healthy {
		t.Error("expected cluster to be healthy")
	}
}
