// Package cli provides the Kavrok command-line interface.
package cli

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/analyzer"
	"github.com/mohammedimrankasab/kavrok/internal/discovery"
	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"github.com/spf13/cobra"
)

func newDoctorCommand(
	clientFactory func() (kubernetes.Client, error),
) *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Diagnose Kubernetes cluster health",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientFactory()
			if err != nil {
				if renderErr := renderDoctorFailure(cmd, err); renderErr != nil {
					return fmt.Errorf(
						"render doctor failure: %w",
						renderErr,
					)
				}

				return err
			}

			snapshot, err := discovery.DiscoverSnapshot(
				cmd.Context(),
				client,
			)
			if err != nil {
				if renderErr := renderDoctorFailure(cmd, err); renderErr != nil {
					return fmt.Errorf(
						"render doctor failure: %w",
						renderErr,
					)
				}

				return err
			}

			findings := analyzer.Analyze(snapshot)

			if err := renderDoctorResult(
				cmd,
				snapshot,
				findings,
			); err != nil {
				return fmt.Errorf(
					"render doctor result: %w",
					err,
				)
			}

			if hasCriticalFindings(findings) {
				return fmt.Errorf("diagnostic checks failed")
			}

			return nil
		},
	}
}

func renderDoctorResult(
	cmd *cobra.Command,
	snapshot discovery.ClusterSnapshot,
	findings []analyzer.Finding,
) error {
	out := cmd.OutOrStdout()

	if _, err := fmt.Fprintln(out, "Kavrok Doctor"); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		out,
		"Cluster:    %s\n",
		snapshot.Cluster.Version,
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		out,
		"Nodes:      %d\n",
		snapshot.Health.NodeCount,
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		out,
		"Namespaces: %d\n",
		len(snapshot.Namespaces),
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		out,
		"Pods:       %d\n",
		snapshot.Workloads.PodCount,
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out, "Findings:"); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	for _, finding := range findings {
		if _, err := fmt.Fprintf(
			out,
			"%s %-18s %s\n",
			severitySymbol(finding.Severity),
			finding.Code,
			finding.Message,
		); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	switch {
	case hasCriticalFindings(findings):
		_, err := fmt.Fprintln(out, "Result: FAILED")
		return err

	case hasWarningFindings(findings):
		_, err := fmt.Fprintln(out, "Result: WARNINGS")
		return err

	default:
		_, err := fmt.Fprintln(out, "Result: PASSED")
		return err
	}
}

func renderDoctorFailure(
	cmd *cobra.Command,
	err error,
) error {
	out := cmd.OutOrStdout()

	if _, writeErr := fmt.Fprintln(out, "Kavrok Doctor"); writeErr != nil {
		return writeErr
	}

	if _, writeErr := fmt.Fprintln(out); writeErr != nil {
		return writeErr
	}

	if _, writeErr := fmt.Fprintf(
		out,
		"✗ kubernetes      Kubernetes unavailable: %v\n",
		err,
	); writeErr != nil {
		return writeErr
	}

	if _, writeErr := fmt.Fprintln(out); writeErr != nil {
		return writeErr
	}

	_, writeErr := fmt.Fprintln(out, "Result: FAILED")

	return writeErr
}

func severitySymbol(
	severity analyzer.Severity,
) string {
	switch severity {
	case analyzer.SeverityCritical:
		return "✗"

	case analyzer.SeverityWarning:
		return "⚠"

	case analyzer.SeverityInfo:
		return "✓"

	default:
		return "?"
	}
}

func hasCriticalFindings(
	findings []analyzer.Finding,
) bool {
	for _, finding := range findings {
		if finding.Severity == analyzer.SeverityCritical {
			return true
		}
	}

	return false
}

func hasWarningFindings(
	findings []analyzer.Finding,
) bool {
	for _, finding := range findings {
		if finding.Severity == analyzer.SeverityWarning {
			return true
		}
	}

	return false
}
