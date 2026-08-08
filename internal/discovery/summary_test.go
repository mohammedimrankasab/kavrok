package discovery

import "testing"

func TestSummarizeCluster(t *testing.T) {
	t.Parallel()

	cluster := ClusterInfo{
		Version:  "v1.34.0",
		Platform: "linux/arm64",
		Commit:   "abc123",
	}

	nodes := NodeSummaryList{
		Nodes: []NodeSummary{
			{
				Name:  "control-plane",
				Ready: true,
			},
			{
				Name:  "worker-1",
				Ready: true,
			},
			{
				Name:  "worker-2",
				Ready: true,
			},
		},
	}

	result := SummarizeCluster(cluster, nodes)

	if result.KubernetesVersion != "v1.34.0" {
		t.Errorf(
			"expected Kubernetes version %q, got %q",
			"v1.34.0",
			result.KubernetesVersion,
		)
	}

	if result.NodeCount != 3 {
		t.Errorf(
			"expected node count %d, got %d",
			3,
			result.NodeCount,
		)
	}

	if result.ReadyNodes != 3 {
		t.Errorf(
			"expected ready nodes %d, got %d",
			3,
			result.ReadyNodes,
		)
	}

	if result.NotReadyNodes != 0 {
		t.Errorf(
			"expected not-ready nodes %d, got %d",
			0,
			result.NotReadyNodes,
		)
	}

	if !result.Healthy {
		t.Error("expected cluster to be healthy")
	}
}

func TestSummarizeClusterWithNotReadyNodes(t *testing.T) {
	t.Parallel()

	cluster := ClusterInfo{
		Version: "v1.34.0",
	}

	nodes := NodeSummaryList{
		Nodes: []NodeSummary{
			{
				Name:  "control-plane",
				Ready: true,
			},
			{
				Name:  "worker-1",
				Ready: false,
			},
			{
				Name:  "worker-2",
				Ready: true,
			},
		},
	}

	result := SummarizeCluster(cluster, nodes)

	if result.NodeCount != 3 {
		t.Errorf("expected 3 nodes, got %d", result.NodeCount)
	}

	if result.ReadyNodes != 2 {
		t.Errorf("expected 2 ready nodes, got %d", result.ReadyNodes)
	}

	if result.NotReadyNodes != 1 {
		t.Errorf("expected 1 not-ready node, got %d", result.NotReadyNodes)
	}

	if result.Healthy {
		t.Error("expected cluster to be unhealthy")
	}
}

func TestSummarizeClusterWithNoNodes(t *testing.T) {
	t.Parallel()

	cluster := ClusterInfo{
		Version: "v1.34.0",
	}

	result := SummarizeCluster(
		cluster,
		NodeSummaryList{},
	)

	if result.NodeCount != 0 {
		t.Errorf("expected 0 nodes, got %d", result.NodeCount)
	}

	if result.ReadyNodes != 0 {
		t.Errorf("expected 0 ready nodes, got %d", result.ReadyNodes)
	}

	if result.NotReadyNodes != 0 {
		t.Errorf("expected 0 not-ready nodes, got %d", result.NotReadyNodes)
	}

	if result.Healthy {
		t.Error("expected empty cluster to be unhealthy")
	}
}
func TestSummarizeClusterHealth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		nodes        NodeSummaryList
		wantTotal    int
		wantReady    int
		wantNotReady int
		wantHealthy  bool
	}{
		{
			name: "all nodes ready",
			nodes: NodeSummaryList{
				Nodes: []NodeSummary{
					{Ready: true},
					{Ready: true},
				},
			},
			wantTotal:    2,
			wantReady:    2,
			wantNotReady: 0,
			wantHealthy:  true,
		},
		{
			name: "one node not ready",
			nodes: NodeSummaryList{
				Nodes: []NodeSummary{
					{Ready: true},
					{Ready: false},
				},
			},
			wantTotal:    2,
			wantReady:    1,
			wantNotReady: 1,
			wantHealthy:  false,
		},
		{
			name:         "no nodes",
			nodes:        NodeSummaryList{},
			wantTotal:    0,
			wantReady:    0,
			wantNotReady: 0,
			wantHealthy:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := SummarizeCluster(
				ClusterInfo{Version: "v1.34.0"},
				tt.nodes,
			)

			if result.NodeCount != tt.wantTotal {
				t.Errorf(
					"expected total %d, got %d",
					tt.wantTotal,
					result.NodeCount,
				)
			}

			if result.ReadyNodes != tt.wantReady {
				t.Errorf(
					"expected ready %d, got %d",
					tt.wantReady,
					result.ReadyNodes,
				)
			}

			if result.NotReadyNodes != tt.wantNotReady {
				t.Errorf(
					"expected not-ready %d, got %d",
					tt.wantNotReady,
					result.NotReadyNodes,
				)
			}

			if result.Healthy != tt.wantHealthy {
				t.Errorf(
					"expected healthy=%v, got %v",
					tt.wantHealthy,
					result.Healthy,
				)
			}
		})
	}
}
