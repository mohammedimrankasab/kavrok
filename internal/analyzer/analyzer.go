package analyzer

import "github.com/mohammedimrankasab/kavrok/internal/discovery"

// Analyze evaluates a Kubernetes snapshot and returns engineering findings.
func Analyze(snapshot discovery.ClusterSnapshot) []Finding {
	findings := make([]Finding, 0)

	findings = append(findings, analyzeNodes(snapshot)...)
	findings = append(findings, analyzePods(snapshot)...)

	if len(findings) == 0 {
		findings = append(findings, Finding{
			Severity: SeverityInfo,
			Code:     "CLUSTER_HEALTHY",
			Title:    "Cluster is healthy",
			Message:  "No known health issues were detected.",
		})
	}

	return findings
}
