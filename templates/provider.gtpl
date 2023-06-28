package salesforce 

import (
)
{{/* the provider unction provides Terraform with an interface to configure your provider and access its resources and datasources */}}
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
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
			"subdomain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_SUBDOMAIN", nil),
			},
			"token_url": {
				Type:        schema.TypeString,
				Required:    true,
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
			{{- range .OperationGroups }}
			"salesforce_{{ humanize .Name | snakize }}": resources.{{ pascalize .Name }}(),
			{{- end }}
		},
		DataSourcesMap: map[string]*schema.Resource{
			{{- range .OperationGroups }}
			"salesforce_{{ humanize .Name | snakize }}": resources.DataResource{{ pascalize .Name }}(),
			{{- end }}
		},
		ConfigureContextFunc: providerConfigure,
	}
}
{{/* if you make your own provider using our repo as a template, you'll likely need to change how yours is configured */}}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

 	clientId := d.Get("client_id").(string)
 	clientSecret := d.Get("client_secret").(string)
 	subdomain := d.Get("subdomain").(string) + ".my.salesforce.com"
	tokenUrl := d.Get("token_url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	config := client.NewConfig()
	config.SetClientId(&clientId)
	config.SetClientSecret(&clientSecret)
	config.SetAccountDomain(&subdomain)
	config.SetTokenUrl(&tokenUrl)
	config.SetUsername(&username)
	config.SetPassword(&password)

	c := client.New(config)

	return c, diags
}
