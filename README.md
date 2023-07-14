# Salesforce Terraform Provider
This provider includes resources from the [Salesforce REST API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/intro_rest.htm). This is intended as a more robust implementation of a Salesforce provider than the [official Salesforce provider](https://registry.terraform.io/providers/hashicorp/salesforce/latest), which is solely focused on user administration. There are many data objects included in the REST API which would benefit from being under infrastructure-as-code, and this provider aims to capture those.

This is a work-in-progress, but any resources already included in this provider are functional. The WIP part is to add new [Tooling API](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/reference_objects_list.htm) models as needed. If you need a resource that is not already included here, please either [submit an issue](https://github.com/rlmartin/terraform-provider-salesforce/issues) or follow [the instructions below](#adding-new-resources) and submit a PR.


## Development

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.14.x
-	[Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)
-   [Go-Swagger](https://goswagger.io/install.html) v0.27.0+ (to generate the code)

### Building the Provider
Clone repository (here, using SSH):
```sh
$ git clone git@github.com:vestahealthcare/salesforce-terraform-provider.git
```
Enter the provider directory and build the provider:
```sh
$ cd salesforce-terraform-provider/
$ make
```
The Makefile will then generate the code, build the binary, and copy it to the Terraform plugin directory.


### Adding new resources
1. Each resource has its own file in [spec_files/components](./spec_files/components). Create a new `{resource_name}.json` file in this directory.
2. Copy one of the existing resource definitions as a template, and paste it into the new `.json` file. A few notes about this definition:
  1. The Terraform provider model requires all four CRUD methods (Create = POST, Request = GET, Update = PATCH/PUT, Delete = DELETE). If you don't define all of these endpoints, the provider will not build.
  2. The CRUD operations for a given resource are unified via a common `tag` on each of the defined paths/methods. An example is: `"tags": ["PlatformEventChannel"]`; this also must be set on the top-level object in the definition.
  3. For the model generation in this repo, each model that you want generated (i.e. all of them) must include `"example": "isResource"`; this flags the model for generation (and is non-obvious).
  4. The supported security option(s) must be defined on each endpoint. For this Salesforce provider this will always be `"security": [{"oauth2_password": []}]`.
  5. The `id` field of the main resource definition (which likely is set, returned, and thus loaded into Terraform state) must _not_ be required.
  6. It is possible that the `id` field _must_ be that name (case-insensitive); I would need to double-check the `go-swagger` code to verify this.
3. Consult the model definitions in the [Tooling API](https://developer.salesforce.com/docs/atlas.en-us.api_tooling.meta/api_tooling/reference_objects_list.htm). Add definitions for your new resource.
4. Replace the copied resource name with your new resource name in all locations.
5. Run `cd spec_files; python3 createSpecFile.py; cd ..`.
6. Run `make build` and ensure it builds successfully.
7. Run `make docs`.
8. Commit all code changes and push a PR.
9. Review/approve/merge the PR.
10. Tag the new version: `git tag vX.Y.X; git push --tags`.
11. The new version will get published as both a [GitHub release](https://github.com/rlmartin/terraform-provider-salesforce/releases) and on the [Terraform Registry](https://registry.terraform.io/providers/rlmartin/salesforce).


### Bootstrapping
The first time deploying this provider will require the following:

1. If you don't already have one, create a [Terraform Registry](https://registry.terraform.io/) account. You will need to grant access to the GitHub org in which the provider is published.
2. Follow the [GitHub documentation for generating a new GPG key](https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key), steps 1-13 (everything except adding it to GitHub, which is not needed).
3. [Add the GPG key](https://registry.terraform.io/settings/gpg-keys/new) to your Terraform Registry account. `Source` and `Source URL` are optional.
4. Export the secret key (`gpg --armor --export-secret-keys {your key id here} | pbcopy`) and add it as a secret on the GitHub repo called `GPG_PRIVATE_KEY`.
5. Add a second secret on the GitHub repo called `PASSPHRASE`, which contains the passphrase you set on the key file.
6. Publish the initial version manually from local:
  1. Create a temporary [GitHub token](https://github.com/settings/tokens) and copy it. It should have `repo` and `write:packages` permissions.
  2. Run `GITHUB_TOKEN={your GitHub token here} GPG_FINGERPRINT={your key id here} goreleaser release --clean`.
  3. Once successful delete the temporary token.
7. [Add the provider](https://registry.terraform.io/publish/provider) in the Terraform Registry. It may take a couple minutes, post-creation, for the first version to appear on the listing.


## Using the provider

The Salesforce Terraform Provider has two methods for setting required arguments:
Environment Variables
```sh
export SALESFORCE_CLIENT_ID=xyz
export SALESFORCE_CLIENT_SECRET=xyz
export SALESFORCE_SUBDOMAIN=xyz
export SALESFORCE_TOKEN_URL=xyz
export SALESFORCE_USERNAME=xyz
export SALESFORCE_PASSWORD=xyz
```

Provider Initialization
```sh
provider "salesforce" {
  client_id     = var.salesforce_client_id
  client_secret = var.salesforce_client_secret
  subdomain     = var.salesforce_subdomain
	token_url     = var.salesforce_token_url
	username      = var.salesforce_username
	password      = var.salesforce_password
}
```
Test cases can be found in the `/salesforce-terraform-provider/Test` directory.


## Source
The basis of this comes from the [LogicMonitor blog post](http://logicmonitor.com/blog/how-to-write-a-custom-terraform-provider-automatically-with-openapi) detailing a Swagger-driven approach to developing a Terraform provider, plus its accompanying [GitHub repo](https://github.com/logicmonitor/automated-terraform-provider)
