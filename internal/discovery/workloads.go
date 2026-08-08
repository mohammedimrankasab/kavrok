package discovery

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// DiscoverWorkloads retrieves pod health information.
func DiscoverWorkloads(
	ctx context.Context,
	client kubernetes.Client,
) (WorkloadSummary, error) {
	pods, err := client.ListPods(ctx)
	if err != nil {
		return WorkloadSummary{}, fmt.Errorf(
			"list Kubernetes pods: %w",
			err,
		)
	}

	result := WorkloadSummary{
		PodCount: len(pods.Items),
	}

	for _, pod := range pods.Items {
		switch pod.Status.Phase {
		case corev1.PodRunning:
			result.RunningPods++

		case corev1.PodPending:
			result.PendingPods++

		case corev1.PodFailed:
			result.FailedPods++
		}

		if podReady(pod) {
			result.ReadyPods++
		}
	}

	result.NotReadyPods = result.PodCount - result.ReadyPods

	return result, nil
}

func podReady(pod corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady {
			return condition.Status == corev1.ConditionTrue
		}
	}

	return false
}
