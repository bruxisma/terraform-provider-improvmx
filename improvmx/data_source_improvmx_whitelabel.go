package improvmx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXWhitelabels() *schema.Resource {
	return &schema.Resource{
		Description: "A list of domains used as whitelabels",
		ReadContext: dataSourceWhitelabelsRead,
		Schema: map[string]*schema.Schema{
			"labels": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceWhitelabelsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	labels, error := session.Account.Labels(ctx)
	if error != nil {
		return diag.Errorf("Error reading ImprovMX whitelabels: %s", error)
	}
	names := make([]string, len(labels))
	for _, label := range labels {
		names = append(names, label.Name)
	}
	data.Set("labels", labels)
	return nil
}
