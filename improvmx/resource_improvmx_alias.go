package improvmx

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

func resourceImprovMXAlias() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceImprovMXAliasCreate,
		ReadContext:   resourceImprovMXAliasRead,
		UpdateContext: resourceImprovMXAliasUpdate,
		DeleteContext: resourceImprovMXAliasDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceImprovMXAliasImport,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: "The ID returned by the ImprovMX API",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"address": &schema.Schema{
				Description: "The email address the alias will forward to",
				Type:        schema.TypeString,
				Required:    true,
			},
			"alias": &schema.Schema{
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

func resourceImprovMXAliasCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	forward := data.Get("address").(string)
	domain := data.Get("domain").(string)
	name := data.Get("alias").(string)
	alias, error := session.Aliases.Create(ctx, domain, name, forward)
	if error != nil {
		return diag.Errorf("ImprovMX: Error creating '%s@%s' for '%s': %s", name, domain, forward, error)
	}
	address := fmt.Sprintf("%s@%s", name, domain)
	data.SetId(address)
	data.Set("id", alias.ID)
	data.Set("address", alias.Address)
	return nil
}

func resourceImprovMXAliasRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)
	name := data.Get("alias").(string)
	alias, error := session.Aliases.Read(ctx, domain, name)
	if error != nil {
		return diag.Errorf("ImprovMX: Error reading alias '%s' from domain '%s': %s", name, domain, error)
	}
	address := fmt.Sprintf("%s@%s", name, domain)
	data.SetId(address)
	data.Set("id", alias.ID)
	data.Set("address", alias.Address)
	return nil
}

func resourceImprovMXAliasUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	forward := data.Get("address").(string)
	domain := data.Get("domain").(string)
	name := data.Get("alias").(string)
	alias, error := session.Aliases.Update(ctx, domain, name, forward)
	if error != nil {
		return diag.Errorf("ImprovMX: Error updating alias '%s' under domain '%s': %s", name, domain, error)
	}
	address := fmt.Sprintf("%s@%s", name, domain)
	data.SetId(address)
	data.Set("id", alias.ID)
	data.Set("address", alias.Address)
	return nil
}

func resourceImprovMXAliasDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	session := meta.(*improvmx.Session)
	domain := data.Get("domain").(string)
	name := data.Get("alias").(string)
	if error := session.Aliases.Delete(ctx, domain, name); error != nil {
		return diag.Errorf("ImprovMX: Error deleting alias '%s' from domain '%s': %s", name, domain, error)
	}
	return nil
}

func resourceImprovMXAliasImport(ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	address := data.Id()
	idx := strings.LastIndex(address, "@")
	if idx < 0 {
		return nil, errors.New(fmt.Sprintf("'%s' is an invalid email address", address))
	}
	name, domain := address[:idx], address[idx+1:]
	data.Set("domain", domain)
	data.Set("alias", name)
	if diagnostics := resourceImprovMXAliasRead(ctx, data, meta); diagnostics != nil {
		return nil, errors.New(fmt.Sprintf("ImprovMX: An error ocurred while importing '%s': %v", address, diagnostics))
	}
	return []*schema.ResourceData{data}, nil
}
