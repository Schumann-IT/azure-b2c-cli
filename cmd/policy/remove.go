package policy

import (
	"github.com/spf13/cobra"
)

var remove = &cobra.Command{
	Use:   "remove",
	Short: "Delete remote b2c policies.",
	Long:  `Delete remote b2c policies from B2C identity experience framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, en, err := getNewRemoteService()
		if err != nil {
			log.Fatalf("failed to delete policies for environment %s: %s", en, err.Error())
		}

		err = s.DeletePolicies()
		if err != nil {
			log.Fatalf("failed to delete policies for environment %s: %s", en, err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(remove)
}
