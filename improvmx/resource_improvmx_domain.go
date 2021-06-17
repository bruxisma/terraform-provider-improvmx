package improvmx

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func resourceImprovMXDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceImprovMXDomainCreate,
		ReadContext:   resourceImprovMXDomainRead,
		UpdateContext: resourceImprovMXDomainUpdate,
		DeleteContext: resourceImprovMXDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceImprovMXDomainImport,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "The name of the domain",
				Type:        schema.TypeString,
				Required:    true,
			},
			"notification_email": &schema.Schema{
				Description: "Email where the notifications related to this domain will be sent",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"whitelabel": &schema.Schema{
				Description: "The parent's domain owner",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"active": &schema.Schema{
				Description: "Whether the domain is currently active",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"added": &schema.Schema{
				Description: "Unix timestamp of when the domain was added",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"display": &schema.Schema{
				Description: "Friendly display name of the domain",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dkim_selector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceImprovMXDomainCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	name := data.Get("name").(string)

	email := data.Get("notification_email").(string)
	label := data.Get("label").(string)
	option := improvmx.DomainOption{Email: email, Label: label}

	domain, error := session.Domains.Create(ctx, name, option)
	if error != nil {
		return diag.Errorf("ImprovMX: Error creating domain '%s': %s", name, error)
	}

	resourceDomainSetFields(data, domain)
	return nil
}

func resourceImprovMXDomainRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	name := data.Get("name").(string)

	domain, error := session.Domains.Read(ctx, name)
	if error != nil {
		return diag.Errorf("ImprovMX: Error reading domain '%s': %s", name, error)
	}
	resourceDomainSetFields(data, domain)
	return nil
}

func resourceImprovMXDomainUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	name := data.Get("name").(string)

	email := data.Get("notification_email").(string)
	label := data.Get("whitelabel").(string)

	option := improvmx.DomainOption{Email: email, Label: label}

	domain, error := session.Domains.Update(ctx, name, option)
	if error != nil {
		return diag.Errorf("ImprovMX: Error updating domain '%s': %s", name, error)
	}
	resourceDomainSetFields(data, domain)
	return nil
}

func resourceImprovMXDomainDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	name := data.Get("name").(string)
	if error := session.Domains.Delete(ctx, name); error != nil {
		return diag.Errorf("ImprovMX: Error deleting domain '%s': %s", name, error)
	}
	return nil
}

func resourceImprovMXDomainImport(ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	name := data.Id()
	data.Set("name", name)
	if diagnostics := resourceImprovMXDomainRead(ctx, data, meta); diagnostics != nil {
		return nil, errors.New(fmt.Sprintf("ImprovMX: An error occurred while import '%s': %v", name, diagnostics))
	}
	return []*schema.ResourceData{data}, nil
}

func resourceDomainSetFields(data *schema.ResourceData, domain *improvmx.Domain) {
	data.SetId(domain.Name)
	data.Set("active", domain.Active)
	data.Set("added", domain.Added.Unix())
	//data.Set("aliases", domain.Aliases)
	data.Set("name", domain.Name)
	data.Set("display", domain.Display)
	data.Set("dkim_selector", domain.DKIMSelector)
	data.Set("notification_email", domain.NotificationEmail)
	data.Set("whitelabel", domain.Whitelabel)
}
