package serviceprincipal

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/uuid"
	"github.com/schumann-it/azure-b2c-sdk-for-go/msgraph"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `Create a cervice-principal to use with CI.`,
	Run: func(cmd *cobra.Command, args []string) {
		dc, _ := azidentity.NewDefaultAzureCredential(nil)
		gc, err := msgraph.NewClientWithCredential(dc)
		if err != nil {
			log.Fatalf("could not create graph client: %s", err.Error())
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("could not parse flag name: %s", err.Error())
		}
		tid, err := cmd.Flags().GetString("tenant")
		if err != nil {
			log.Fatalf("could not parse flag tenant: %s", err.Error())
		}
		_, err = uuid.Parse(tid)
		if err != nil {
			log.Fatalf("tenant must be a valid uuid, got %s: %s", tid, err.Error())
		}

		if gc.FindServicePrincipal(name) {
			log.Infof("service principal %s exists.", name)
			return
		}

		// create service principal with required graph api permissions
		appId, password, err := gc.CreateServicePrincipal(name, "00000003-0000-0000-c000-000000000000", []string{
			"246dd0d5-5bd0-4def-940b-0421030a5b68", // https://learn.microsoft.com/en-us/graph/permissions-reference#policyreadall
			"79a677f7-b79d-40d0-a36a-3e6f8688dd7a", // https://learn.microsoft.com/en-us/graph/permissions-reference#policyreadwritetrustframework
			"fff194f1-7dce-4428-8301-1badb5518201", // https://learn.microsoft.com/en-us/graph/permissions-reference#trustframeworkkeysetreadall
			"4a771c9a-1cf2-4609-b88e-3d3e02d539cd", // https://learn.microsoft.com/en-us/graph/permissions-reference#trustframeworkkeysetreadwriteall
		})
		if err != nil {
			log.Fatalf("failed to create service principal: %s", err.Error())
		}

		log.Infof("service principal %s created", to.String(appId))
		log.Infof("to use this service principal for deployment, set the following environment valiables: \nexport B2C_ARM_CLIENT_ID=%s\nexport B2C_ARM_CLIENT_SECRET=%s\nexport B2C_ARM_TENANT_ID=%s", to.String(appId), to.String(password), tid)
		log.Info("Ready! Do not forget to grant Admin Consent on the created permissions!")
	},
}

func init() {
	create.Flags().String("tenant", "t", "The tenant ID.")
	create.Flags().StringP("name", "n", "sp-azureb2c-cli", "The name of the service principal to create.")
	RootCmd.AddCommand(create)
}
