package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Glimmer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Glimmer v0.1.0")
		},
	}
}
