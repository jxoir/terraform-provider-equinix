package equinix

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"EQUINIX_API_ID",
				}, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"equinix_ecx_connection": resourceConnection(),
		},
	}
}
