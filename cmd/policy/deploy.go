package policy

import (
	"github.com/hashicorp/go-multierror"
	b2c "github.com/schumann-it/azure-b2c-sdk-for-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var deploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy b2c policy.",
	Long:  `Deploy b2c policy to B2C identity experience framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, en, err := getDeployService(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to deploy policies for environment %s: %s", en, err.Error())
		}

		err = s.DeployPolicies(en)
		if err != nil {
			log.Fatalf("failed to deploy policies for environment %s: %s", en, err.Error())
		}
	},
}

func init() {
	deploy.Flags().StringP("build-dir", "b", "./build", "Build directory")
	RootCmd.AddCommand(deploy)
}

func getDeployService(flags *pflag.FlagSet) (*b2c.Service, string, error) {
	var errs error

	cf, en, err := getConfigAndEnvironment()
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	bd, err := getAbsPath(flags.GetString("build-dir"))
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	s, err := b2c.NewServiceFromConfigFile(cf)
	s.MustWithTargetDir(bd)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	return s, en, errs
}
