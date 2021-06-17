package improvmx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXSMTPCredentials() *schema.Resource {
	return &schema.Resource{
		Description: "All credentials for the given domain",
		ReadContext: dataSourceImprovMXSMTPCredentialsRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: "The domain to query for users",
				Type:        schema.TypeString,
				Required:    true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created": &schema.Schema{
							Description: "Unix timestamp when the credential was created",
							Computed:    true,
							Type:        schema.TypeInt,
						},
						"username": &schema.Schema{
							Description: "The username for this credential instance",
							Computed:    true,
							Type:        schema.TypeString,
						},
						"usage": &schema.Schema{
							Description: "Usage numbers for the account",
							Computed:    true,
							Type:        schema.TypeInt,
						},
					},
				},
			},
		},
	}
}

func dataSourceImprovMXSMTPCredentialsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)

	list, error := session.Credentials.List(ctx, domain)

	if error != nil {
		return diag.Errorf("ImprovMX: Error getting SMTP credentials for domain '%s': %s", domain, error)
	}
	credentialIDs := make([]string, 0, len(list))
	credentials := make([]map[string]interface{}, len(list))

	for _, credential := range list {
		credentials = append(credentials, map[string]interface{}{
			"username": credential.Username,
			"created":  credential.CreatedAt.Unix(),
			"usage":    credential.Usage,
		})
	}

	data.SetId(generateID(credentialIDs))
	data.Set("users", credentials)
	return nil
}
