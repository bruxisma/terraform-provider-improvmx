package improvmx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func resourceImprovMXSMTPCredential() *schema.Resource {
	return &schema.Resource{
		Description:   `ImprovMX SMTP Credentials. NOTE: These are a premium feature`,
		CreateContext: resourceImprovMXSMTPCredentialCreate,
		UpdateContext: resourceImprovMXSMTPCredentialUpdate,
		DeleteContext: resourceImprovMXSMTPCredentialDelete,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: "The name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": &schema.Schema{
				Description: "Username for the SMTP login",
				Type:        schema.TypeString,
				Required:    true,
			},
			"password": &schema.Schema{
				Description: "Password for the SMTP login",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"created": &schema.Schema{
				Description: "The time the SMTP credential was created",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceImprovMXSMTPCredentialCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)

	username := data.Get("username").(string)
	password := data.Get("password").(string)

	login := improvmx.User{Username: username, Password: password}

	credential, error := session.Credentials.Create(ctx, domain, login)
	if error != nil {
		return diag.Errorf("ImprovMX: Error creating credential '%s' for domain '%s': %s", username, domain, error)
	}

	data.SetId(credential.Username)
	data.Set("created", credential.CreatedAt.Unix())

	return nil
}

func resourceImprovMXSMTPCredentialUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)

	username := data.Get("username").(string)
	password := data.Get("password").(string)

	login := improvmx.User{Username: username, Password: password}

	credential, error := session.Credentials.Update(ctx, domain, login)
	if error != nil {
		return diag.Errorf("ImprovMX: Error updating credential '%s' for domain '%s': %s", username, domain, error)
	}

	data.SetId(credential.Username)
	data.Set("created", credential.CreatedAt.Unix())

	return nil
}

func resourceImprovMXSMTPCredentialDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)

	username := data.Get("username").(string)

	if error := session.Credentials.Delete(ctx, domain, username); error != nil {
		return diag.Errorf("ImprovMX: Error deleting credential '%s' for domain '%s': %s", username, domain, error)
	}

	return nil
}
