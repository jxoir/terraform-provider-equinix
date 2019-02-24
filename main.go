package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jxoir/terraform-provider-equinix/equinix"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: equinix.Provider})

}
