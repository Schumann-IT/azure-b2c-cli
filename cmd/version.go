package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"
)

var printVersion = &cobra.Command{
	Use:   "version",
	Short: "Print version.",
	Long:  `Print version.`,
	Run: func(cmd *cobra.Command, args []string) {
		println(version)
	},
}

func init() {
	rootCmd.AddCommand(printVersion)
}
