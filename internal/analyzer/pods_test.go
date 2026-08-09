package analyzer

import (
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
	corev1 "k8s.io/api/core/v1"
)

func TestAnalyzePodsUnschedulableIncludesResourceRequest(t *testing.T) {
	t.Parallel()

	snapshot := discovery.ClusterSnapshot{
		Nodes: discovery.NodeSummaryList{
			Nodes: []discovery.NodeSummary{
				{
					Name:              "test-node",
					AllocatableMemory: "7Gi",
				},
			},
		},
		Workloads: discovery.WorkloadSummary{
			Pods: []discovery.PodSummary{
				{
					Name:            "kavrok-pending",
					Namespace:       "default",
					Phase:           "Pending",
					PendingReason:   "Unschedulable",
					PendingMessage:  "0/1 nodes are available: 1 Insufficient memory.",
					RequestedMemory: "100Gi",
				},
			},
		},
	}

	findings := Analyze(snapshot)

	var finding *Finding

	for i := range findings {
		if findings[i].Code == "POD_UNSCHEDULABLE" {
			finding = &findings[i]
			break
		}
	}

	if finding == nil {
		t.Fatal("expected POD_UNSCHEDULABLE finding")
	}

	expectedEvidence := map[string]string{
		"node":               "test-node",
		"requested_memory":   "100Gi",
		"allocatable_memory": "7Gi",
		"diagnosis":          "requested memory exceeds node allocatable memory",
	}

	for key, expected := range expectedEvidence {
		found := false

		for _, evidence := range finding.Evidence {
			if evidence.Key == key && evidence.Value == expected {
				found = true
				break
			}
		}

		if !found {
			t.Errorf(
				"expected evidence %s=%q",
				key,
				expected,
			)
		}
	}
}

func TestAnalyzePodsUnschedulableIncludesCPUResourceRequest(t *testing.T) {
	t.Parallel()

	snapshot := discovery.ClusterSnapshot{
		Nodes: discovery.NodeSummaryList{
			Nodes: []discovery.NodeSummary{
				{
					Name:           "test-node",
					AllocatableCPU: "2",
				},
			},
		},
		Workloads: discovery.WorkloadSummary{
			PendingPods: 1,
			Pods: []discovery.PodSummary{
				{
					Name:           "cpu-heavy",
					Namespace:      "default",
					Phase:          string(corev1.PodPending),
					PendingReason:  "Unschedulable",
					PendingMessage: "0/1 nodes are available: Insufficient cpu.",
					RequestedCPU:   "4",
				},
			},
		},
	}

	findings := analyzePods(snapshot)

	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}

	var (
		requestedCPU   string
		allocatableCPU string
		node           string
		diagnosis      string
	)

	for _, evidence := range findings[0].Evidence {
		switch evidence.Key {
		case "requested_cpu":
			requestedCPU = evidence.Value
		case "allocatable_cpu":
			allocatableCPU = evidence.Value
		case "node":
			node = evidence.Value
		case "diagnosis":
			diagnosis = evidence.Value
		}
	}

	if requestedCPU != "4" {
		t.Fatalf(
			"expected requested_cpu evidence with value 4, got %q",
			requestedCPU,
		)
	}

	if allocatableCPU != "2" {
		t.Fatalf(
			"expected allocatable_cpu evidence with value 2, got %q",
			allocatableCPU,
		)
	}

	if node != "test-node" {
		t.Fatalf(
			"expected node evidence with value test-node, got %q",
			node,
		)
	}

	if diagnosis != "requested CPU exceeds node allocatable CPU" {
		t.Fatalf(
			"unexpected diagnosis: %q",
			diagnosis,
		)
	}
}
