package policy

import (
	"github.com/hashicorp/go-multierror"
	"github.com/schumann-it/azure-b2c-sdk-for-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "Build",
	Long:  `Build source policy and replacing template variables for given environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, en, err := getBuildService(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to build policies for environment %s: %s", en, err.Error())
		}

		err = s.BuildPolicies(en)
		if err != nil {
			log.Fatalf("failed to build policies for environment %s: %s", en, err.Error())
		}
	},
}

func init() {
	build.Flags().StringP("source", "s", "./src", "File directory")
	build.Flags().StringP("destination", "d", "./build", "Destination directory")
	RootCmd.AddCommand(build)
}

func getBuildService(flags *pflag.FlagSet) (*b2c.Service, string, error) {
	var errs error

	cf, en, err := getConfigAndEnvironment()
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	sd, err := getAbsPath(flags.GetString("source"))
	if err != nil {
		errs = multierror.Append(errs, err)
	}
	dd, err := getAbsPath(flags.GetString("destination"))
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	s, err := b2c.NewServiceFromConfigFile(cf)
	s.MustWithSourceDir(sd)
	s.MustWithTargetDir(dd)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	return s, en, errs
}
