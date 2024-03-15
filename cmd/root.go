package cmd

import (
	"com.schumann-it.azure-b2c-cli/cmd/policy"
	"github.com/spf13/cobra"
	"gopkg.in/op/go-logging.v1"
)

var log = logging.MustGetLogger("cmd")

var rootCmd = &cobra.Command{
	Use:   "azureb2c",
	Short: "Tooling for Azure B2C",
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		lvl := logging.INFO
		isDebug, _ := cmd.PersistentFlags().GetBool("debug")
		if isDebug {
			lvl = logging.DEBUG
		}
		logging.SetLevel(lvl, "")
		logging.SetFormatter(logging.MustStringFormatter(
			`%{color}%{level}(%{module})%{color:reset} %{message}`,
		))
		log = logging.MustGetLogger("cmd")
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	rootCmd.AddCommand(policy.RootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
