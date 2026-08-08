package analyzer

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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

		message := fmt.Sprintf(
			"%s/%s: %s",
			pod.Namespace,
			pod.Name,
			pod.PendingMessage,
		)

		finding := Finding{
			Severity: SeverityWarning,
			Code:     "POD_UNSCHEDULABLE",
			Title:    "Pod cannot be scheduled",
			Message:  message,
		}

		if evidence := memorySchedulingEvidence(
			snapshot,
			pod.RequestedMemory,
		); len(evidence) > 0 {
			finding.Evidence = append(
				finding.Evidence,
				evidence...,
			)
		}

		findings = append(findings, finding)
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

func memorySchedulingEvidence(
	snapshot discovery.ClusterSnapshot,
	requestedMemory string,
) []Evidence {
	requested, err := resource.ParseQuantity(requestedMemory)
	if err != nil {
		return nil
	}

	for _, node := range snapshot.Nodes.Nodes {
		if node.AllocatableMemory == "" {
			continue
		}

		allocatable, err := resource.ParseQuantity(
			node.AllocatableMemory,
		)
		if err != nil {
			continue
		}

		if requested.Cmp(allocatable) <= 0 {
			return nil
		}

		return []Evidence{
			{
				Key:   "node",
				Value: node.Name,
			},
			{
				Key:   "requested_memory",
				Value: requestedMemory,
			},
			{
				Key:   "allocatable_memory",
				Value: node.AllocatableMemory,
			},
			{
				Key:   "diagnosis",
				Value: "requested memory exceeds node allocatable memory",
			},
		}
	}

	return nil
}
