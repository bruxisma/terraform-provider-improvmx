package improvmx

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"occult.work/improvmx"
)

// This is based off of (but not a direct copy of!) how cloudflare combines ids
func generateID(ids []string) string {
	sort.Strings(ids)
	hash := sha256.New()
	hash.Write([]byte(strings.Join(ids, "")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func domainSetFields(data *schema.ResourceData, domain *improvmx.Domain) {
	data.SetId(domain.Name)
	data.Set("active", domain.Active)
	data.Set("added", domain.Added.Unix())
	data.Set("name", domain.Name)
	data.Set("display", domain.Display)
	data.Set("dkim_selector", domain.DKIMSelector)
	data.Set("notification_email", domain.NotificationEmail)
	data.Set("whitelabel", domain.Whitelabel)
}
