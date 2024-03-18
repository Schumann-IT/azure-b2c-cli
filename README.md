# azure-b2c-cli

A tool for automating [Azure B2C](https://learn.microsoft.com/en-us/azure/active-directory-b2c/overview).

**Features:**
* [x] Build and Deploy Custom Policies
* [] Validate certain infrastructure components

The project has been inspired by
* [go-ieftool](https://github.com/judedaryl/go-ieftool)
* [VS Code extension](https://github.com/azure-ad-b2c/vscode-extension)

The [VS Code extension](https://github.com/azure-ad-b2c/vscode-extension) can replace variables within custom policies 
to prepare them for deployment to different environments. But it comes with a few limitations
* it cannot be used in CI systems
* it only supports a flat directory structure

[go-ieftool](https://github.com/judedaryl/go-ieftool) has resolved the issues with VS Code, but the creation of
required resources, such as [Policy Keys](https://learn.microsoft.com/en-us/azure/active-directory-b2c/policy-keys-overview?pivots=b2c-custom-policy)
have still to be managed manually.

Please also see the 
* [Terraform Provider](https://registry.terraform.io/providers/Schumann-IT/azureadb2c/latest)
* [B2C SDK](https://github.com/Schumann-IT/azure-b2c-sdk-for-go)

## Install

### Dev Installation

Either clone this repo and build it
```
go build -o azure-b2c-cli 
```

or use Go to install directly into your $GOBIN directory (e.g. $GOPATH/bin):
```
go install github.com/Schumann-IT/azure-b2c-cli 
```

## Usage

You can use `--help` on each subcommand to get help. Get started with
```
azure-b2c-cli --help
```

You can generate completion scripts
```
azure-b2c-cli completion [bash|zsh|fish|powershell]
```

## Policy Management

### Build

* Download the [Custom Policy Starter Pack](https://github.com/Azure-Samples/active-directory-b2c-custom-policy-starterpack)
and replace all occurrences of `yourtenant.onmicrosoft.com` with `{Settings:Tenant}`.
* Create a `config.yaml` in one of the subdirectories, eg. `LocalAccounts`
```yaml
- name: stage
  settings:
    Tenant: <your-stage-tenant>.onmicrosoft.com    
- name: prod
  settings:
    Tenant: <your-prod-tenant>.onmicrosoft.com    
```
* Run 
```
azure-b2c-cli policy build --config ./config.yaml --environment stage --source ./ --destination ./build 
azure-b2c-cli policy build --config ./config.yaml --environment prod --source ./ --destination ./build 
```
* Find the deployable policies in `build`

### Deploy

To be able to deploy policies, some resources need to be configured. The `Starter Pack`, you have downloaded before, 
requires 
* [signing and encryption keys](https://learn.microsoft.com/en-us/azure/active-directory-b2c/tutorial-create-user-flows?pivots=b2c-custom-policy#add-signing-and-encryption-keys-for-identity-experience-framework-applications) 
* [Identity Experiance Framework Apps](https://learn.microsoft.com/en-us/azure/active-directory-b2c/tutorial-create-user-flows?pivots=b2c-custom-policy#add-signing-and-encryption-keys-for-identity-experience-framework-applications)

Either follow the documentation or tryout the [Terraform Provider](https://registry.terraform.io/providers/Schumann-IT/azureadb2c/latest)

* Deploy
```
# The service principal requires the following permissions:
#  * Policy.Read.All
#  * Policy.ReadWrite.TrustFramework

export B2C_ARM_CLIENT_ID=<stage-service-principal-app-id>
export B2C_ARM_CLIENT_SECRET=<stage-service-principal-app-password>
export B2C_ARM_TENANT_ID=<stage-tenant-id>

azure-b2c-cli policy deploy --config ./config.yaml --environment stage --build-dir ./build 
```