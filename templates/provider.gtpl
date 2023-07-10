package salesforce 

import (
)
{{/* the provider unction provides Terraform with an interface to configure your provider and access its resources and datasources */}}
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"subdomain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SF_SUBDOMAIN", nil),
			},
			{{- range .SecurityDefinitions }}
			{{- if and .IsOAuth2 (eq .Flow "password") }}
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
			{{ end }}
			{{ end -}}
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

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := client.NewConfig()

 	domain := d.Get("subdomain").(string) + ".my.salesforce.com"
	config.SetAccountDomain(&domain)
	{{- range .SecurityDefinitions }}
	{{- if and .IsOAuth2 (eq .Flow "password") }}
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
	{{ end -}}
	{{ end }}

	c := client.New(config)

	return c, diags
}
