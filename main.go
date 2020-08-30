package main

import (
	"github.com/halvardssm/terraform-provider-domeneshop/domeneshop"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: domeneshop.Provider,
	})
}
