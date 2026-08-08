package analyzer

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
	corev1 "k8s.io/api/core/v1"
)

func analyzePods(
	snapshot discovery.ClusterSnapshot,
) []Finding {
	findings := make([]Finding, 0)

	workloads := snapshot.Workloads

	unschedulablePods := 0

	for _, pod := range workloads.Pods {
		if pod.Phase != string(corev1.PodPending) {
			continue
		}

		if pod.PendingReason != "Unschedulable" &&
			pod.PendingReason != "FailedScheduling" {
			continue
		}

		unschedulablePods++

		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "POD_UNSCHEDULABLE",
			Title:    "Pod cannot be scheduled",
			Message: fmt.Sprintf(
				"%s/%s: %s",
				pod.Namespace,
				pod.Name,
				pod.PendingMessage,
			),
		})
	}

	genericPendingPods := workloads.PendingPods - unschedulablePods

	if genericPendingPods > 0 {
		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "PODS_PENDING",
			Title:    "Pods are Pending",
			Message: fmt.Sprintf(
				"%d pods are Pending.",
				genericPendingPods,
			),
		})
	}

	if workloads.FailedPods > 0 {
		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "PODS_FAILED",
			Title:    "Pods have failed",
			Message: fmt.Sprintf(
				"%d pods are in Failed state.",
				workloads.FailedPods,
			),
		})
	}

	if workloads.NotReadyPods > 0 {
		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "PODS_NOT_READY",
			Title:    "Pods are not Ready",
			Message: fmt.Sprintf(
				"%d pods are not Ready.",
				workloads.NotReadyPods,
			),
		})
	}

	return findings
}
