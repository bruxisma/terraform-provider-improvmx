package improvmx

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXAliases() *schema.Resource {
	return &schema.Resource{
		Description: "All aliases under a given domain",
		ReadContext: dataSourceImprovMXAliasesRead,
		Schema: map[string]*schema.Schema{
			"aliases": &schema.Schema{
				Description: "The set of aliases returned",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward": &schema.Schema{
							Description: "The email address the alias will forward to",
							Computed:    true,
							Type:        schema.TypeString,
						},
						"alias": &schema.Schema{
							Description: "The name of the alias",
							Computed:    true,
							Type:        schema.TypeString,
						},
						"id": &schema.Schema{
							Description: "The ID returned by the ImprovMX API",
							Computed:    true,
							Type:        schema.TypeInt,
						},
					},
				},
			},
			"filter": &schema.Schema{
				Description: "Options to filter the results",
				Optional:    true,
				Type:        schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"starts_with": &schema.Schema{
							Description: "Aliases that start with the given string",
							Optional:    true,
							Type:        schema.TypeString,
						},
						"is_active": &schema.Schema{
							Description: "Aliases for domains that are active",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"page": &schema.Schema{
							Description: "Page to start from",
							Optional:    true,
							Type:        schema.TypeInt,
						},
					},
				},
			},
			"domain": &schema.Schema{
				Description: "The domain to look up",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceImprovMXAliasesRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)
	list, error := session.Aliases.List(ctx, domain)
	if error != nil {
		return diag.Errorf("ImprovMX: Error getting aliases for domain '%s': %s", domain, error)
	}
	aliasIds := make([]string, 0, len(list))
	aliases := make([]map[string]interface{}, len(list))
	for _, alias := range list {
		aliases = append(aliases, map[string]interface{}{
			"forward": alias.Address,
			"alias":   alias.Name,
			"id":      alias.ID,
		})
		aliasIds = append(aliasIds, fmt.Sprintf("%s@%s", alias.Name, domain))
	}

	data.SetId(generateID(aliasIds))
	data.Set("aliases", aliases)
	return nil
}
