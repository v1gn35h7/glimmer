package worker

import "github.com/spf13/cobra"

func NewWorkerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Start the Glimmer worker",
		Run: func(cmd *cobra.Command, args []string) {
			// Worker logic goes here
		},
	}
}
