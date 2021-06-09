package improvmx

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXAlias() *schema.Resource {
	return &schema.Resource{
		Description: "An email alias for a given domain",
		ReadContext: dataSourceImprovMXAccountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "The ID returned by the ImprovMX API",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"address": &schema.Schema{
				Description: "The email address the alias will forward to",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": &schema.Schema{
				Description: "The name of the alias",
				Type:        schema.TypeString,
				Required:    true,
			},
			"domain": &schema.Schema{
				Description: "The domain for the alias",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceImprovMXAliasRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)
	name := data.Get("name").(string)
	alias, error := session.Aliases.Read(ctx, domain, name)
	if error != nil {
		return diag.Errorf("Error reading alias '%s' from domain '%s': %s", error)
	}
	data.SetId(strconv.FormatInt(alias.ID, 10))
	data.Set("address", alias.Address)
	return nil
}
