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
