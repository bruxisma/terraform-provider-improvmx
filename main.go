package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/ixm-one/terraform-provider-improvmx/improvmx"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: improvmx.Provider,
	})
}
