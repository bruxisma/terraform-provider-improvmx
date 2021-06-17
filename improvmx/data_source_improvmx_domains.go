package improvmx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXDomains() *schema.Resource {
	return &schema.Resource{
		Description: "All domains for the authenticated ImprovMX account",
		ReadContext: dataSourceImprovMXDomainsRead,
		Schema: map[string]*schema.Schema{
			"domains": &schema.Schema{
				Description: "All domain objects returned",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"notification_email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dkim_selector": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"whitelabel": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"added": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"filter": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceImprovMXDomainsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	results, error := session.Domains.List(ctx)
	if error != nil {
		return diag.Errorf("ImprovMX: Error getting domains: %s", error)
	}
	domainIds := make([]string, 0, len(results))
	domains := make([]map[string]interface{}, len(results))
	for _, domain := range results {
		domains = append(domains, map[string]interface{}{
			"active":             domain.Active,
			"name":               domain.Name,
			"notification_email": domain.NotificationEmail,
			"display":            domain.Display,
			"dkim_selector":      domain.DKIMSelector,
			"whitelabel":         domain.Whitelabel,
			"added":              domain.Added.Unix(),
		})
		domainIds = append(domainIds, domain.Name)
	}

	data.SetId(generateID(domainIds))
	data.Set("domains", domains)
	return nil
}
