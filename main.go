package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"occult.work/terraform-provider-improvmx/improvmx"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: improvmx.Provider,
	})
}
