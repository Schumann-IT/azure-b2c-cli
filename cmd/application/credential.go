package application

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/schumann-it/azure-b2c-sdk-for-go/msgraph"
	"github.com/spf13/cobra"
)

var credential = &cobra.Command{
	Use:   "credential",
	Short: "credential",
	Long:  `Add an application password credential`,
	Run: func(cmd *cobra.Command, args []string) {
		dc, _ := azidentity.NewDefaultAzureCredential(nil)
		gc, err := msgraph.NewClientWithCredential(dc)
		if err != nil {
			log.Fatalf("could not create graph client: %s", err.Error())
		}

		name, _ := cmd.Flags().GetString("name")
		oid, _ := cmd.Flags().GetString("oid")
		c, err := gc.AddApplicationPasswordCredentials(oid, name)
		if err != nil {
			log.Fatalf("could not create password credential for %s: %s", oid, err.Error())
		}

		log.Infof("CLIENT_SECRET: %s created for application %s", to.String(c.GetSecretText()), oid)
	},
}

func init() {
	credential.Flags().String("name", "", "The name of the credential.")
	_ = credential.MarkFlagRequired("name")
	credential.Flags().String("oid", "", "The application object id.")
	_ = credential.MarkFlagRequired("oid")
	RootCmd.AddCommand(credential)
}
