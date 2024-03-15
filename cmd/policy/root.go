package policy

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	b2c "github.com/schumann-it/azure-b2c-sdk-for-go"
	"github.com/spf13/cobra"
	"gopkg.in/op/go-logging.v1"
)

var log = logging.MustGetLogger("azureb2c/policy")

var RootCmd = &cobra.Command{
	Use:   "policy",
	Short: "Tooling for Azure B2C policy management",
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "./config.yaml", "Path to the configuration file")
	RootCmd.PersistentFlags().StringP("environment", "e", "", "Environment to deploy (default: all environments)")
}

func getConfigAndEnvironment() (string, string, error) {
	var errs error

	cf, err := RootCmd.Flags().GetString("config")
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	en, err := RootCmd.Flags().GetString("environment")
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	if en == "" {
		errs = multierror.Append(errs, fmt.Errorf("must provide flag 'environment'; got: %s", en))
	}

	return cf, en, errs
}

func getAbsPath(p string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if !filepath.IsAbs(p) {
		p, err = filepath.Abs(p)
		if err != nil {
			return "", err
		}
	}

	return p, nil
}

func getNewRemoteService() (*b2c.Service, string, error) {
	var errs error

	cf, en, err := getConfigAndEnvironment()
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	s, err := b2c.NewService(cf, "", "")
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	return s, en, errs
}
