package analyzer

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
)

func analyzePods(
	snapshot discovery.ClusterSnapshot,
) []Finding {
	findings := make([]Finding, 0)

	workloads := snapshot.Workloads

	if workloads.PendingPods > 0 {
		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "PODS_PENDING",
			Title:    "Pods are Pending",
			Message: fmt.Sprintf(
				"%d pods are Pending.",
				workloads.PendingPods,
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
