package discovery

import (
	"context"
	"errors"
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
)

type fakeWorkloadClient struct {
	pods *corev1.PodList
	err  error
}

func (f *fakeWorkloadClient) ServerVersion() (*version.Info, error) {
	return &version.Info{
		GitVersion: "v1.36.1",
	}, nil
}

func (f *fakeWorkloadClient) ClusterName() string {
	return "test-cluster"
}

func (f *fakeWorkloadClient) ListNodes(
	_ context.Context,
) (*corev1.NodeList, error) {
	return &corev1.NodeList{}, nil
}

func (f *fakeWorkloadClient) ListNamespaces(
	_ context.Context,
) (*corev1.NamespaceList, error) {
	return &corev1.NamespaceList{}, nil
}

func (f *fakeWorkloadClient) ListPods(
	_ context.Context,
) (*corev1.PodList, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.pods, nil
}

var _ kubernetes.Client = (*fakeWorkloadClient)(nil)

func TestDiscoverWorkloads(t *testing.T) {
	t.Parallel()

	client := &fakeWorkloadClient{
		pods: &corev1.PodList{
			Items: []corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "running-ready",
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
						Name:      "running-not-ready",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
						Conditions: []corev1.PodCondition{
							{
								Type:   corev1.PodReady,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pending",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodPending,
						Conditions: []corev1.PodCondition{
							{
								Type:    corev1.PodScheduled,
								Status:  corev1.ConditionFalse,
								Reason:  "Unschedulable",
								Message: "Insufficient memory",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "failed",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodFailed,
					},
				},
			},
		},
	}

	result, err := DiscoverWorkloads(context.Background(), client)
	if err != nil {
		t.Fatalf("DiscoverWorkloads() failed: %v", err)
	}

	if result.PodCount != 4 {
		t.Errorf("expected PodCount=4, got %d", result.PodCount)
	}

	if result.RunningPods != 2 {
		t.Errorf("expected RunningPods=2, got %d", result.RunningPods)
	}

	if result.PendingPods != 1 {
		t.Errorf("expected PendingPods=1, got %d", result.PendingPods)
	}

	if result.FailedPods != 1 {
		t.Errorf("expected FailedPods=1, got %d", result.FailedPods)
	}

	if result.ReadyPods != 1 {
		t.Errorf("expected ReadyPods=1, got %d", result.ReadyPods)
	}

	if result.NotReadyPods != 1 {
		t.Errorf("expected NotReadyPods=1, got %d", result.NotReadyPods)
	}

	if len(result.Pods) != 4 {
		t.Fatalf("expected 4 pod summaries, got %d", len(result.Pods))
	}

}

func TestDiscoverWorkloadsPendingPodIsNotNotReady(t *testing.T) {
	t.Parallel()

	client := &fakeWorkloadClient{
		pods: &corev1.PodList{
			Items: []corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "unschedulable",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodPending,
						Conditions: []corev1.PodCondition{
							{
								Type:    corev1.PodScheduled,
								Status:  corev1.ConditionFalse,
								Reason:  "Unschedulable",
								Message: "0/1 nodes are available: Insufficient memory",
							},
						},
					},
				},
			},
		},
	}

	result, err := DiscoverWorkloads(context.Background(), client)
	if err != nil {
		t.Fatalf("DiscoverWorkloads() failed: %v", err)
	}

	if result.PendingPods != 1 {
		t.Fatalf(
			"expected PendingPods=1, got %d",
			result.PendingPods,
		)
	}

	if result.NotReadyPods != 0 {
		t.Fatalf(
			"expected NotReadyPods=0 for a pending pod, got %d",
			result.NotReadyPods,
		)
	}

	if len(result.Pods) != 1 {
		t.Fatalf(
			"expected 1 pod summary, got %d",
			len(result.Pods),
		)
	}

	pod := result.Pods[0]

	if pod.Name != "unschedulable" {
		t.Errorf(
			"expected pod name %q, got %q",
			"unschedulable",
			pod.Name,
		)
	}

	if pod.PendingReason != "Unschedulable" {
		t.Errorf(
			"expected PendingReason=%q, got %q",
			"Unschedulable",
			pod.PendingReason,
		)
	}

	if pod.PendingMessage != "0/1 nodes are available: Insufficient memory" {
		t.Errorf(
			"unexpected PendingMessage: %q",
			pod.PendingMessage,
		)
	}

}

func TestDiscoverWorkloadsListPodsFailure(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("connection refused")

	client := &fakeWorkloadClient{
		err: expectedErr,
	}

	_, err := DiscoverWorkloads(context.Background(), client)
	if err == nil {
		t.Fatal("expected DiscoverWorkloads() to fail")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected error to wrap %q, got %v",
			expectedErr,
			err,
		)
	}

}

func TestPodReady(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		conditions []corev1.PodCondition
		expected   bool
	}{
		{
			name: "ready",
			conditions: []corev1.PodCondition{
				{
					Type:   corev1.PodReady,
					Status: corev1.ConditionTrue,
				},
			},
			expected: true,
		},
		{
			name: "not ready",
			conditions: []corev1.PodCondition{
				{
					Type:   corev1.PodReady,
					Status: corev1.ConditionFalse,
				},
			},
			expected: false,
		},
		{
			name:       "condition missing",
			conditions: nil,
			expected:   false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pod := corev1.Pod{
				Status: corev1.PodStatus{
					Conditions: tt.conditions,
				},
			}

			if got := podReady(pod); got != tt.expected {
				t.Errorf(
					"podReady()=%v, expected %v",
					got,
					tt.expected,
				)
			}
		})
	}

}

func TestPodPendingCondition(t *testing.T) {
	t.Parallel()

	pod := corev1.Pod{
		Status: corev1.PodStatus{
			Conditions: []corev1.PodCondition{
				{
					Type:    corev1.PodScheduled,
					Status:  corev1.ConditionFalse,
					Reason:  "Unschedulable",
					Message: "Insufficient memory",
				},
			},
		},
	}

	condition := podPendingCondition(pod)
	if condition == nil {
		t.Fatal("expected pending condition")
	}

	if condition.Reason != "Unschedulable" {
		t.Errorf(
			"expected reason %q, got %q",
			"Unschedulable",
			condition.Reason,
		)
	}

	if condition.Message != "Insufficient memory" {
		t.Errorf(
			"expected message %q, got %q",
			"Insufficient memory",
			condition.Message,
		)
	}

}

func TestPodPendingConditionNotFound(t *testing.T) {
	t.Parallel()

	pod := corev1.Pod{
		Status: corev1.PodStatus{
			Conditions: []corev1.PodCondition{
				{
					Type:   corev1.PodReady,
					Status: corev1.ConditionFalse,
				},
			},
		},
	}

	if condition := podPendingCondition(pod); condition != nil {
		t.Fatal("expected no pending condition")
	}

}
