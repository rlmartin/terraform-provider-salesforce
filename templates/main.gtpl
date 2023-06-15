package main

import (
	"vestahealthcare/salesforce"
)
{{/* Every TF provider needs a main function. Look how little it is, innit cute? */}}
func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return salesforce.Provider()
		},
	}

	plugin.Serve(opts)
}
