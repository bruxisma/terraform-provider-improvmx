package improvmx

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Description: "The API token used to connect to ImprovMX.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_TOKEN", nil),
			},
			"base_url": {
				Description: "The base URL for ImprovMX's API. Must end in a /",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_BASE_URL", "https://api.improvmx.com/v3/"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"improvmx_alias":  resourceImprovMXAlias(),
			"improvmx_domain": resourceImprovMXDomain(),
			//	"improvmx_smtp_credentials": resourceImprovMXSMTPCredentials(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"improvmx_account":     dataSourceImprovMXAccount(),
			"improvmx_whitelabels": dataSourceImprovMXWhitelabels(),
			"improvmx_domains":     dataSourceImprovMXDomains(),
			"improvmx_aliases":     dataSourceImprovMXAliases(),
			"improvmx_domain":      dataSourceImprovMXDomain(),
			"improvmx_alias":       dataSourceImprovMXAlias(),
		},
	}
	provider.ConfigureContextFunc = func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		session, error := improvmx.New(
			data.Get("token").(string),
			improvmx.WithBaseURL(data.Get("base_url").(string)),
			improvmx.WithUserAgent(fmt.Sprintf("terraform-provider-improvmx/%s", provider.TerraformVersion)))
		if error != nil {
			return nil, diag.FromErr(error)
		}
		return session, nil
	}
	return provider
}
