package analyzer

import (
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
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
