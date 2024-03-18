package serviceprincipal

import (
	"github.com/spf13/cobra"
	"gopkg.in/op/go-logging.v1"
)

var log = logging.MustGetLogger("azureb2c/serviceprincipal")

var RootCmd = &cobra.Command{
	Use:   "serviceprincipal",
	Short: "Tooling for Azure B2C service principals",
}
