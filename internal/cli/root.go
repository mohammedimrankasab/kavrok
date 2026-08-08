// Package cli provides the Kavrok command-line interface.
package cli

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"github.com/spf13/cobra"
)

// Dependencies contains the dependencies required by CLI commands.
type Dependencies struct {
	KubernetesClientFactory func() (kubernetes.Client, error)
}

// NewRootCommand creates the root Kavrok command.
func NewRootCommand(deps Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kavrok",
		Short:         "Kubernetes engineering intelligence CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newVersionCommand(),
		newDoctorCommand(deps.KubernetesClientFactory),
	)

	return cmd
}

// Execute runs the Kavrok CLI.
func Execute() error {
	kubeClient, err := kubernetes.New()
	if err != nil {
		return fmt.Errorf("create kubernetes client: %w", err)
	}
	Deps := Dependencies{
		KubernetesClientFactory: func() (kubernetes.Client, error) {
			return kubeClient, nil
		},
	}
	return NewRootCommand(Deps).Execute()
}
