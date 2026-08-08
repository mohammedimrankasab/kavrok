package discovery

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// DiscoverNodes retrieves node information from the Kubernetes cluster.
func DiscoverNodes(
	ctx context.Context,
	client kubernetes.Client,
) (NodeSummaryList, error) {
	nodes, err := client.ListNodes(ctx)
	if err != nil {
		return NodeSummaryList{}, fmt.Errorf(
			"list Kubernetes nodes: %w",
			err,
		)
	}

	result := NodeSummaryList{
		Nodes: make([]NodeSummary, 0, len(nodes.Items)),
	}

	for _, node := range nodes.Items {
		result.Nodes = append(
			result.Nodes,
			nodeSummary(node),
		)
	}

	return result, nil
}

func nodeSummary(node corev1.Node) NodeSummary {
	return NodeSummary{
		Name:              node.Name,
		Ready:             nodeReady(node),
		Roles:             nodeRoles(node),
		KubernetesVersion: node.Status.NodeInfo.KubeletVersion,
		OS:                node.Status.NodeInfo.OperatingSystem,
		Architecture:      node.Status.NodeInfo.Architecture,
		AllocatableCPU:    node.Status.Allocatable.Cpu().String(),
		AllocatableMemory: node.Status.Allocatable.Memory().String(),
	}
}

func nodeReady(node corev1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady {
			return condition.Status == corev1.ConditionTrue
		}
	}

	return false
}

func nodeRoles(node corev1.Node) []string {
	roles := make([]string, 0)

	for label := range node.Labels {
		const prefix = "node-role.kubernetes.io/"

		if len(label) > len(prefix) && label[:len(prefix)] == prefix {
			role := label[len(prefix):]

			if role != "" {
				roles = append(roles, role)
			}
		}
	}

	return roles
}
