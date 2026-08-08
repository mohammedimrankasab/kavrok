package analyzer

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/discovery"
)

func analyzeNodes(
	snapshot discovery.ClusterSnapshot,
) []Finding {
	findings := make([]Finding, 0)

	if snapshot.Health.NodeCount == 0 {
		return append(findings, Finding{
			Severity: SeverityCritical,
			Code:     "NO_NODES",
			Title:    "Cluster has no nodes",
			Message:  "No Kubernetes nodes were discovered.",
		})
	}

	if snapshot.Health.NotReadyNodes > 0 {
		findings = append(findings, Finding{
			Severity: SeverityWarning,
			Code:     "NODE_NOT_READY",
			Title:    "Nodes are not Ready",
			Message: fmt.Sprintf(
				"%d of %d nodes are not Ready.",
				snapshot.Health.NotReadyNodes,
				snapshot.Health.NodeCount,
			),
		})
	}

	return findings
}
