package oozie

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"oozie_server": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OOZIE_SERVER", nil),
				Description: "The host and port of Oozie server.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
	}
}
