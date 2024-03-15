package policy

import (
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List remote b2c policy.",
	Long:  `List remote b2c policy from B2C identity experience framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, en, err := getNewRemoteService()
		if err != nil {
			log.Fatalf("could list policies for environment %s: %s", en, err.Error())
		}

		err = s.ListPolicies()
		if err != nil {
			log.Fatalf("could list policies for environment %s: %s", en, err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(list)
}
