package main

import (
	"vestahealthcare/salesforce"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return salesforce.Provider()
		},
	}

	plugin.Serve(opts)
}
