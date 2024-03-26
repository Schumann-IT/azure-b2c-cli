package application

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/schumann-it/azure-b2c-sdk-for-go/msgraph"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `Create an empty application`,
	Run: func(cmd *cobra.Command, args []string) {
		dc, _ := azidentity.NewDefaultAzureCredential(nil)
		gc, err := msgraph.NewClientWithCredential(dc)
		if err != nil {
			log.Fatalf("could not create graph client: %s", err.Error())
		}

		names, _ := cmd.Flags().GetStringArray("name")
		for _, n := range names {
			a, err := gc.CreateApplication(n)
			if err != nil {
				log.Errorf("could not create app %s: %s", n, err.Error())
				continue
			}
			log.Infof("application %s created with objectId=%s, clientId=%s", n, to.String(a.GetId()), to.String(a.GetAppId()))
		}
	},
}

func init() {
	create.Flags().StringArray("name", []string{}, "The name of the application to create (can be supplied multiple times).")
	_ = create.MarkFlagRequired("name")
	RootCmd.AddCommand(create)
}
