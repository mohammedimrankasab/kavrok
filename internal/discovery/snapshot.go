package discovery

import (
	"context"
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// DiscoverSnapshot retrieves the current Kubernetes cluster state.
func DiscoverSnapshot(
	ctx context.Context,
	client kubernetes.Client,
) (ClusterSnapshot, error) {
	cluster, err := Discover(client)
	if err != nil {
		return ClusterSnapshot{}, fmt.Errorf(
			"discover cluster: %w",
			err,
		)
	}

	nodes, err := DiscoverNodes(ctx, client)
	if err != nil {
		return ClusterSnapshot{}, fmt.Errorf(
			"discover nodes: %w",
			err,
		)
	}

	namespaces, err := DiscoverNamespaces(ctx, client)
	if err != nil {
		return ClusterSnapshot{}, fmt.Errorf(
			"discover namespaces: %w",
			err,
		)
	}

	workloads, err := DiscoverWorkloads(ctx, client)
	if err != nil {
		return ClusterSnapshot{}, fmt.Errorf(
			"discover workloads: %w",
			err,
		)
	}

	health := SummarizeCluster(cluster, nodes)

	return ClusterSnapshot{
		Cluster:    cluster,
		Nodes:      nodes,
		Namespaces: namespaces,
		Workloads:  workloads,
		Health:     health,
	}, nil
}
