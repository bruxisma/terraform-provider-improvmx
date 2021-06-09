package improvmx

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func dataSourceImprovMXAccount() *schema.Resource {
	return &schema.Resource{
		Description: "Details about your ImprovMX account",
		ReadContext: dataSourceImprovMXAccountRead,
		Schema: map[string]*schema.Schema{
			"billing_email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"premium": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceImprovMXAccountRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	account, error := session.Account.Read(ctx)
	if error != nil {
		return diag.Errorf("Error reading from ImprovMX Account endpoint: %s", error)
	}
	hash := sha256.Sum256([]byte(account.Email))
	data.SetId(hex.EncodeToString(hash[:]))
	data.Set("billing_email", account.BillingEmail)
	data.Set("password", account.Password)
	data.Set("premium", account.Premium)
	data.Set("email", account.Email)
	return nil
}
