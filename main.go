package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tmihalicka/terraform-provider-oozie/oozie"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: oozie.Provider})
}
