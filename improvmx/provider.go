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
				Type:        schema.TypeString,
				Optional:    false,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_TOKEN", nil),
				Description: "The API token used to connect to ImprovMX.",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_BASE_URL", "https://api.improvmx.com/v3/"),
				Description: "The base URL for ImprovMX's API. Must end in a /",
			},
		},

		//ResourcesMap: map[string]*schema.Resource{
		//	"improvmx_alias":            resourceImprovMXAlias(),
		//	"improvmx_domain":           resourceImprovMXDomain(),
		//	"improvmx_smtp_credentials": resourceImprovMXSMTPCredentials(),
		//},

		DataSourcesMap: map[string]*schema.Resource{
			"improvmx_account": dataSourceImprovMXAccount(),
			"improvmx_alias":   dataSourceImprovMXAlias(),
			//"improvmx_alias_logs":  dataSourceImprovMXAliasLogs(),
			//"improvmx_domain": dataSourceImprovMXDomain(),
			//"improvmx_domain_logs": dataSourceImprovMXDomainLogs(),
			"improvmx_whitelabels": dataSourceImprovMXWhitelabels(),
		},
	}
	provider.ConfigureContextFunc = func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		session, error := improvmx.New(
			data.Get("token").(string),
			improvmx.WithHostURL(data.Get("base_url").(string)),
			improvmx.WithUserAgent(fmt.Sprintf("terraform-provider-improvmx/%s", provider.TerraformVersion)))
		if error != nil {
			return nil, diag.FromErr(error)
		}
		return session, nil
	}
	return provider
}
