package coordinator

import "github.com/spf13/cobra"

func NewCoordinatorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "coordinator",
		Short: "Start the Glimmer coordinator",
		Run: func(cmd *cobra.Command, args []string) {
			// Coordinator logic goes here
		},
	}
}
