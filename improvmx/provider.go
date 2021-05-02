package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/slurps-mad-rips/improvmx-go"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    false,
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

		ResourcesMap: map[string]*schema.Resource{
			"improvmx_alias":            resourceImprovMXAlias(),
			"improvmx_domain":           resourceImprovMXDomain(),
			"improvmx_smtp_credentials": resourceImprovMXSMTPCredentials(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"improvmx_account":     dataSourceImprovMXAccount(),
			"improvmx_alias":       dataSourceImprovMXAlias(),
			"improvmx_alias_logs":  dataSourceImprovMXAliasLogs(),
			"improvmx_domain":      dataSourceImprovMXDomain(),
			"improvmx_domain_logs": dataSourceImprovMXDomainLogs(),
			"improvmx_whitelabel":  dataSourceImprovMXWhitelabel(),
		},
	}

	//p.ConfigureFunc = providerConfigure(p)

	return p
}

func configure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   data.Get("token").(string),
		BaseURL: data.Get("base_url").(string),
	}
	meta, err := config.Meta()
	if err != nil {
		return nil, err
	}
	meta.(*Owner).StopContext = p.StopContext()
	return meta, nil
}
