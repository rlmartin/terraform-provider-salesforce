package main

import (
	"vestahealthcare/salesforce"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --rendered-provider-name=Salesforce

func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return salesforce.Provider()
		},
	}

	plugin.Serve(opts)
}
