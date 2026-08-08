package discovery

// SummarizeCluster creates a high-level health summary from discovered data.
func SummarizeCluster(
	cluster ClusterInfo,
	nodes NodeSummaryList,
) ClusterSummary {
	readyNodes := 0

	for _, node := range nodes.Nodes {
		if node.Ready {
			readyNodes++
		}
	}

	nodeCount := len(nodes.Nodes)
	notReadyNodes := nodeCount - readyNodes

	return ClusterSummary{
		KubernetesVersion: cluster.Version,
		NodeCount:         nodeCount,
		ReadyNodes:        readyNodes,
		NotReadyNodes:     notReadyNodes,
		Healthy:           nodeCount > 0 && notReadyNodes == 0,
	}
}
