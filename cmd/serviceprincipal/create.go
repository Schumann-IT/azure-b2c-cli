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
	Long:  `Create a service principal to use with CI.`,
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

		gc.CreateOrganizationClient(tid)
		o, err := gc.OrganizationClient.Get()
		if err != nil {
			log.Fatalf("could not read tenant information: %s", err.Error())
		}
		tn := o.GetDisplayName()

		if gc.FindServicePrincipal(name) {
			log.Infof("service principal %s exists for tenant '%s'.", name, to.String(tn))
			return
		}

		// create service principal with required graph api permissions
		appId, password, err := gc.CreateServicePrincipal(name, "00000003-0000-0000-c000-000000000000", []string{
			// read and write application registrations
			"9a5d68dd-52b0-4cc2-bd40-abcf44ac3a30", // https://learn.microsoft.com/en-us/graph/permissions-reference#applicationreadall
			"1bfefb4e-e0b5-418b-a88f-73c46d2cc8e9", // https://learn.microsoft.com/en-us/graph/permissions-reference#applicationreadwriteall
			// read and write organization information
			"498476ce-e0fe-48b0-b801-37ba7e2685c6", // https://learn.microsoft.com/en-us/graph/permissions-reference#organizationreadall
			"292d869f-3427-49a8-9dab-8c70152b74e9", // https://learn.microsoft.com/en-us/graph/permissions-reference#organizationreadwriteall
			// read and write organization branding
			"eb76ac34-0d62-4454-b97c-185e4250dc20", // https://learn.microsoft.com/en-us/graph/permissions-reference#organizationalbrandingreadall
			"d2ebfbc1-a5f8-424b-83a6-56ab5927a73c", // https://learn.microsoft.com/en-us/graph/permissions-reference#organizationalbrandingreadwriteall
			// read and write policies
			"246dd0d5-5bd0-4def-940b-0421030a5b68", // https://learn.microsoft.com/en-us/graph/permissions-reference#policyreadall
			"79a677f7-b79d-40d0-a36a-3e6f8688dd7a", // https://learn.microsoft.com/en-us/graph/permissions-reference#policyreadwritetrustframework
			"fff194f1-7dce-4428-8301-1badb5518201", // https://learn.microsoft.com/en-us/graph/permissions-reference#trustframeworkkeysetreadall
			"4a771c9a-1cf2-4609-b88e-3d3e02d539cd", // https://learn.microsoft.com/en-us/graph/permissions-reference#trustframeworkkeysetreadwriteall
			// read tenant information
			"cac88765-0581-4025-9725-5ebc13f729ee", // https://learn.microsoft.com/en-us/graph/permissions-reference#crosstenantinformationreadbasicall
		})
		if err != nil {
			log.Fatalf("failed to create service principal for tenant '%s': %s", to.String(tn), err.Error())
		}

		log.Infof("service principal %s created for tenant '%s'", to.String(appId), to.String(tn))
		log.Infof("to use this service principal for deployment, set the following environment valiables: \nexport B2C_ARM_CLIENT_ID=%s\nexport B2C_ARM_CLIENT_SECRET=%s\nexport B2C_ARM_TENANT_ID=%s", to.String(appId), to.String(password), tid)
		log.Info("Ready! Do not forget to grant Admin Consent on the created permissions!")
	},
}

func init() {
	create.Flags().String("tenant", "", "The tenant ID.")
	_ = create.MarkFlagRequired("tenant")
	create.Flags().StringP("name", "n", "sp-azureb2c-cli", "The name of the service principal to create.")
	RootCmd.AddCommand(create)
}
