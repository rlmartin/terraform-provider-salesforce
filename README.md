# Salesforce Terraform Provider


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
