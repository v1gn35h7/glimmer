package cmd

import (
	"fmt"
	"os"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
	"github.com/v1gn35h7/glimmer/pkg/cmd/coordinator"
	"github.com/v1gn35h7/glimmer/pkg/cmd/version"
	"github.com/v1gn35h7/glimmer/pkg/cmd/worker"
)

var ascii = figlet4go.NewAsciiRender()

var rootCmd = &cobra.Command{
	Use:   "glimmer",
	Long:  "Glimmer is a blazing-fast, Go-native stream processing engine that empowers developers to write distributed, stateful stream jobs using idiomatic Go.",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	// add subcommands
	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(coordinator.NewCoordinatorCmd())
	rootCmd.AddCommand(worker.NewWorkerCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}

	renderStr, _ := ascii.RenderOpts("Glimmer", options)
	fmt.Print(renderStr)
}
