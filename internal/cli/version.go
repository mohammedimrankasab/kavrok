package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mohammedimrankasab/kavrok/internal/version"
)

func newVersionCommand() *cobra.Command {
	var short bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display Kavrok version information",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			info := version.Get()

			if short {
				_, err := fmt.Fprintln(cmd.OutOrStdout(), info.Version)
				return err
			}

			_, err := fmt.Fprintf(
				cmd.OutOrStdout(),
				"Kavrok %s\n"+
					"Commit:     %s\n"+
					"Build Date: %s\n"+
					"Go Version: %s\n"+
					"Platform:   %s\n"+
					"Tree State: %s\n",
				info.Version,
				info.Commit,
				info.BuildDate,
				info.GoVersion,
				info.Platform,
				info.TreeState,
			)

			return err
		},
	}
	cmd.Flags().BoolVar(
		&short,
		"short",
		false,
		"Display only the version",
	)

	return cmd
}
