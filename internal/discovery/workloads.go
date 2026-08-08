package discovery

import (
	"context"
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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
		Pods:     make([]PodSummary, 0, len(pods.Items)),
	}

	for _, pod := range pods.Items {
		podSummary := PodSummary{
			Name:            pod.Name,
			Namespace:       pod.Namespace,
			Phase:           string(pod.Status.Phase),
			Ready:           podReady(pod),
			RequestedCPU:    podRequestedResource(pod, corev1.ResourceCPU),
			RequestedMemory: podRequestedResource(pod, corev1.ResourceMemory),
		}

		if pod.Status.Phase == corev1.PodPending {
			if condition := podPendingCondition(pod); condition != nil {
				podSummary.PendingReason = condition.Reason
				podSummary.PendingMessage = condition.Message
			}
		}
		result.Pods = append(result.Pods, podSummary)

		switch pod.Status.Phase {
		case corev1.PodRunning:
			result.RunningPods++

			if podReady(pod) {
				result.ReadyPods++
			}

		case corev1.PodPending:
			result.PendingPods++

		case corev1.PodFailed:
			result.FailedPods++
		}
	}

	result.NotReadyPods = result.RunningPods - result.ReadyPods

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
func podPendingCondition(
	pod corev1.Pod,
) *corev1.PodCondition {
	for i := range pod.Status.Conditions {
		condition := &pod.Status.Conditions[i]

		if condition.Type == corev1.PodScheduled &&
			condition.Status == corev1.ConditionFalse {
			return condition
		}
	}

	return nil
}

func podRequestedResource(
	pod corev1.Pod,
	resourceName corev1.ResourceName,
) string {
	total := resource.Quantity{}

	for _, container := range pod.Spec.Containers {
		if quantity, ok := container.Resources.Requests[resourceName]; ok {
			total.Add(quantity)
		}
	}

	if total.IsZero() {
		return ""
	}

	return total.String()
}
