package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mohammedimrankasab/kavrok/internal/doctor"
	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

func newDoctorCommand(
	kubeClientFactory func() (kubernetes.Client, error),
) *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Diagnose the Kubernetes environment",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			result := doctor.New(kubeClientFactory).Run()

			if err := doctor.WriteReport(
				cmd.OutOrStdout(),
				result,
			); err != nil {
				return err
			}

			if !result.Passed() {
				return fmt.Errorf("diagnostic checks failed")
			}

			return nil
		},
	}
}
