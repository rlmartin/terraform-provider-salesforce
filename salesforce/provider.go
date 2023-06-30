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
			"subdomain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_SUBDOMAIN", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_CLIENT_SECRET", nil),
			},
			"token_url": {
				Type:        schema.TypeString,
				Required:    false,
				DefaultFunc: schema.EnvDefaultFunc("SF_TOKEN_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_PASSWORD", nil),
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

	config := client.NewConfig()

	domain := d.Get("subdomain").(string) + ".my.salesforce.com"
	config.SetAccountDomain(&domain)
	clientId := d.Get("client_id").(string)
	config.SetClientId(&clientId)
	clientSecret := d.Get("client_secret").(string)
	config.SetClientSecret(&clientSecret)
	tokenUrl := d.Get("token_url").(string)
	config.SetTokenUrl(&tokenUrl)
	username := d.Get("username").(string)
	config.SetUsername(&username)
	password := d.Get("password").(string)
	config.SetPassword(&password)

	c := client.New(config)

	return c, diags
}
