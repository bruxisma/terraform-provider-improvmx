package improvmx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceImprovMXDomainRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "The name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"notification_email": &schema.Schema{
				Description: "Email where the notifications related to this domain will be sent",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"whitelabel": &schema.Schema{
				Description: "The parent's domain owner",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"active": &schema.Schema{
				Description: "Whether the domain is currently active",
				Computed:    true,
				Type:        schema.TypeBool,
			},
			"added": &schema.Schema{
				Description: "Unix timestamp of when the domain was added",
				Computed:    true,
				Type:        schema.TypeInt,
			},
			"display": &schema.Schema{
				Description: "Friendly display name of the domain",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"dkim_selector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceImprovMXDomainRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	name := data.Get("name").(string)

	domain, error := session.Domains.Read(ctx, name)
	if error != nil {
		return diag.Errorf("ImprovMX: Error reading from domain '%s': %s", name, error)
	}
	domainSetFields(data, domain)
	return nil
}
