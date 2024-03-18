package application

import (
	"github.com/spf13/cobra"
	"gopkg.in/op/go-logging.v1"
)

var log = logging.MustGetLogger("azureb2c/application")

var RootCmd = &cobra.Command{
	Use:   "application",
	Short: "Tooling for Azure B2C application management",
}
