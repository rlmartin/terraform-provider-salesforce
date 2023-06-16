package salesforce

import (
	"context"
	"vestahealthcare/client"
	"vestahealthcare/salesforce/resources"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_ID", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_KEY", nil),
			},
			"company": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_COMPANY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"salesforce_platform_event_channel": resources.PlatformEventChannel(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"salesforce_platform_event_channel": resources.DataResourcePlatformEventChannel(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	id := d.Get("api_id").(string)
	key := d.Get("api_key").(string)
	company := d.Get("company").(string) + ".salesforce.com"
	config := client.NewConfig()
	config.SetAccessKey(&key)
	config.SetAccessID(&id)
	config.SetAccountDomain(&company)

	c := client.New(config)

	return c, diags
}
