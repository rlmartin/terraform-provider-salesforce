package main

import (
	"vestahealthcare/salesforce"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --rendered-provider-name=Salesforce

{{/* Every TF provider needs a main function. Look how little it is, innit cute? */}}
func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return salesforce.Provider()
		},
	}

	plugin.Serve(opts)
}
