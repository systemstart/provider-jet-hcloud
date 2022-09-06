package network

import (
	"github.com/crossplane/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_id"] = config.Reference{
			Type:      "Network",
			Extractor: "github.com/crossplane-contrib/provider-jet-hcloud/apis/common.ExtractResourceID()",
		}
		// well, ...
		// r.TerraformResource.Schema["network_id"].Type = schema.TypeString
	})
}
