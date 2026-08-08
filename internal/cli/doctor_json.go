package cli

import (
	"encoding/json"
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/analyzer"
	"github.com/mohammedimrankasab/kavrok/internal/discovery"
	"github.com/spf13/cobra"
)

type doctorJSON struct {
	Cluster  doctorJSONCluster   `json:"cluster"`
	Health   doctorJSONHealth    `json:"health"`
	Findings []doctorJSONFinding `json:"findings"`
	Result   string              `json:"result"`
}

type doctorJSONCluster struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type doctorJSONHealth struct {
	Nodes      int `json:"nodes"`
	Namespaces int `json:"namespaces"`
	Pods       int `json:"pods"`
}

type doctorJSONFinding struct {
	Severity string            `json:"severity"`
	Code     string            `json:"code"`
	Title    string            `json:"title"`
	Message  string            `json:"message"`
	Evidence map[string]string `json:"evidence,omitempty"`
}

func renderDoctorJSON(
	cmd *cobra.Command,
	snapshot discovery.ClusterSnapshot,
	findings []analyzer.Finding,
) error {
	result := "PASSED"

	switch {
	case hasCriticalFindings(findings):
		result = "FAILED"

	case hasWarningFindings(findings):
		result = "WARNINGS"
	}

	output := doctorJSON{
		Cluster: doctorJSONCluster{
			Name:    snapshot.Cluster.Name,
			Version: snapshot.Cluster.Version,
		},
		Health: doctorJSONHealth{
			Nodes:      snapshot.Health.NodeCount,
			Namespaces: len(snapshot.Namespaces),
			Pods:       snapshot.Workloads.PodCount,
		},
		Findings: make([]doctorJSONFinding, 0, len(findings)),
		Result:   result,
	}

	for _, finding := range findings {
		output.Findings = append(
			output.Findings,
			doctorJSONFinding{
				Severity: string(finding.Severity),
				Code:     finding.Code,
				Title:    finding.Title,
				Message:  finding.Message,
				Evidence: findingEvidence(finding),
			},
		)
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("encode doctor result: %w", err)
	}

	_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))

	return err
}
