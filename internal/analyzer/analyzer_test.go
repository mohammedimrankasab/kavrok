package analyzer

import (
	"testing"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
)

func TestAnalyzeHealthyCluster(t *testing.T) {
	t.Parallel()

	snapshot := discovery.ClusterSnapshot{
		Cluster: discovery.ClusterInfo{
			Version: "v1.34.0",
		},
		Health: discovery.ClusterSummary{
			KubernetesVersion: "v1.34.0",
			NodeCount:         3,
			ReadyNodes:        3,
			NotReadyNodes:     0,
			Healthy:           true,
		},
		Workloads: discovery.WorkloadSummary{
			PodCount:     10,
			RunningPods:  10,
			ReadyPods:    10,
			NotReadyPods: 0,
		},
	}

	findings := Analyze(snapshot)

	if len(findings) != 1 {
		t.Fatalf(
			"expected 1 finding, got %d",
			len(findings),
		)
	}

	if findings[0].Code != "CLUSTER_HEALTHY" {
		t.Fatalf(
			"expected CLUSTER_HEALTHY, got %q",
			findings[0].Code,
		)
	}

	if findings[0].Severity != SeverityInfo {
		t.Fatalf(
			"expected info severity, got %q",
			findings[0].Severity,
		)
	}
}

func TestAnalyzeUnhealthyCluster(t *testing.T) {
	t.Parallel()

	snapshot := discovery.ClusterSnapshot{
		Health: discovery.ClusterSummary{
			NodeCount:     3,
			ReadyNodes:    2,
			NotReadyNodes: 1,
			Healthy:       false,
		},
		Workloads: discovery.WorkloadSummary{
			PodCount:     10,
			PendingPods:  2,
			FailedPods:   1,
			ReadyPods:    6,
			NotReadyPods: 4,
		},
	}

	findings := Analyze(snapshot)

	if len(findings) != 4 {
		t.Fatalf(
			"expected 4 findings, got %d",
			len(findings),
		)
	}

	expected := []string{
		"NODE_NOT_READY",
		"PODS_PENDING",
		"PODS_FAILED",
		"PODS_NOT_READY",
	}

	for i, code := range expected {
		if findings[i].Code != code {
			t.Errorf(
				"expected finding %d to be %q, got %q",
				i,
				code,
				findings[i].Code,
			)
		}
	}
}

func TestAnalyzeEmptyCluster(t *testing.T) {
	t.Parallel()

	snapshot := discovery.ClusterSnapshot{
		Health: discovery.ClusterSummary{
			NodeCount: 0,
		},
	}

	findings := Analyze(snapshot)

	if len(findings) != 1 {
		t.Fatalf(
			"expected 1 finding, got %d",
			len(findings),
		)
	}

	if findings[0].Code != "NO_NODES" {
		t.Fatalf(
			"expected NO_NODES, got %q",
			findings[0].Code,
		)
	}

	if findings[0].Severity != SeverityCritical {
		t.Fatalf(
			"expected critical severity, got %q",
			findings[0].Severity,
		)
	}
}
